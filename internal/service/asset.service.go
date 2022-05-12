package service

import (
	"fmt"
	"github.com/spacesedan/profile-tracker/internal/datastores"
	"github.com/spacesedan/profile-tracker/internal/dto"
	"github.com/spacesedan/profile-tracker/internal/models"
	"github.com/spacesedan/profile-tracker/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"os"
	"sort"
	"time"
)

type AssetService interface {
	GetAssets(url, slug, walletAddress string) models.Assets
	Producer(ch chan<- *models.TaskSingleAsset, task *models.TaskAssets)
	AssetConsumer(chA <-chan *models.TaskSingleAsset) []*models.AssetEntity
	TraitsWithFloor(traitMap map[string]*models.DBTraits, token *models.DBToken) []*models.DBTraits
	ScrapedAssetsConsumer(chA <-chan *models.TaskSingleAsset) []*models.AssetEntity
	HandleAssets(request dto.AssetRequest) []*models.AssetEntity
	GetRarestTrait(asset *models.AssetEntity) *models.DBTraits
	GetTokenIds(assets []*models.AssetEntity) []string
	GetOwnedTokenIds(req dto.AssetsWithCursorRequest) []string
	GetAssetsRecursively(assets []*models.AssetEntity, req dto.AssetsWithCursorRequest) []*models.AssetEntity
	GetStats(url, slug string) models.OSStats
}

type assetService struct {
	dao *datastores.DAO
}

func NewAssetService(dao *datastores.DAO) AssetService {
	return &assetService{
		dao: dao,
	}
}

// GetAssetsRecursively Keep making request to the OpenSea API until the next cursor is nil or ""
func (a *assetService) GetAssetsRecursively(assets []*models.AssetEntity, req dto.AssetsWithCursorRequest) []*models.AssetEntity {
	var res models.Assets

	var uri string

	uri = os.Getenv("FIRE_PROX") + "assets?owner=" + req.Owner + "&collection_slug=" + req.Slug + "&order_direction=asc&limit=50&cursor=" + req.Cursor
	utils.MakeOpenSeaRequest(uri, &res)

	assets = append(assets, res.Assets...)
	if res.Next == "" {
		return assets
	} else {
		// If the cursor is not "" call the function again
		return a.GetAssetsRecursively(assets, dto.AssetsWithCursorRequest{
			Slug:   req.Slug,
			Owner:  req.Owner,
			Cursor: res.Next,
		})
	}

}

func (a *assetService) GetAssets(url, slug, walletAddress string) models.Assets {
	var assets models.Assets
	uri := url + "assets?owner=" + walletAddress + "&collection=" + slug + "&order_by=pk&order_direction=desc&limit=50"
	utils.GetJson(uri, &assets)
	log.Println(assets.Previous, assets.Previous)
	return assets
}

func (a *assetService) GetStats(url, slug string) models.OSStats {
	uri := url + "collection/" + slug + "/stats"
	var stats models.OSStats
	utils.GetJson(uri, &stats)

	return stats
}

func (a *assetService) GetOwnedTokenIds(req dto.AssetsWithCursorRequest) []string {
	var assets []*models.AssetEntity

	assets = a.GetAssetsRecursively(assets, req)
	if assets == nil {
		for i := 0; i < 10; i++ {
			assets = a.GetAssetsRecursively(assets, req)
			if assets != nil {
				break
			}
		}
	}
	return a.GetTokenIds(assets)
}

func (a *assetService) GetRarestTrait(asset *models.AssetEntity) *models.DBTraits {
	sort.Slice(asset.Traits, func(i, j int) bool {
		return asset.Traits[i].TraitCount > asset.Traits[j].TraitCount
	})

	var rarest = &models.DBTraits{}

	rarest.ID.TraitType = asset.Traits[0].TraitType
	rarest.ID.Value = asset.Traits[0].Value

	return rarest
}

func (a *assetService) Producer(ch chan<- *models.TaskSingleAsset, task *models.TaskAssets) {
	for _, asset := range task.Assets {
		var single *models.TaskSingleAsset
		single = &models.TaskSingleAsset{
			Asset:      asset,
			Collection: task.Collection,
			FloorPrice: task.FloorPrice,
			Traits:     task.Traits,
		}
		ch <- single
	}
}

func (a *assetService) AssetConsumer(chA <-chan *models.TaskSingleAsset) []*models.AssetEntity {
	var owned []*models.AssetEntity

	for {
		select {
		case msg := <-chA:
			if len(msg.Asset.Traits) != 0 {
				rarestTrait := a.GetRarestTrait(msg.Asset)
				msg.Asset.TopTrait = rarestTrait
			}
			fmt.Printf("%+v\n", msg.Asset.TopTrait)
			msg.Asset.FloorPrice = msg.FloorPrice
			msg.Asset.TraitFloorPrice = msg.FloorPrice
			msg.Asset.TopTrait.RarityScore = 0
			msg.Asset.TopTrait.FloorPrice.Price = 0
			msg.Asset.Rank = 0
			msg.Asset.TopTrait.FloorPrice.PriceEntryTime = primitive.NewDateTimeFromTime(time.Now())
			msg.Asset.DBName = "null"
			owned = append(owned, msg.Asset)
		case <-time.After(50 * time.Millisecond):
			return owned
		}
	}
}

func (a *assetService) ScrapedAssetsConsumer(chA <-chan *models.TaskSingleAsset) []*models.AssetEntity {
	var owned []*models.AssetEntity

	for {
		select {
		case msg := <-chA:
			token := a.dao.Repo.NewCollectionQuery().GetToken(msg.Collection, msg.Asset.TokenID)
			traitsWithValue := a.TraitsWithFloor(msg.Traits, token)
			msg.Asset.FloorPrice = msg.FloorPrice
			msg.Asset.TraitFloorPrice = traitsWithValue[0].FloorPrice.Price
			msg.Asset.TopTrait = traitsWithValue[0]
			msg.Asset.DBName = msg.Collection
			msg.Asset.Rank = token.RarityScoreRank
			owned = append(owned, msg.Asset)
		case <-time.After(50 * time.Millisecond):
			return owned
		}
	}
}

func (a *assetService) TraitsWithFloor(traitMap map[string]*models.DBTraits, token *models.DBToken) []*models.DBTraits {
	var withFloor []*models.DBTraits

	for _, attribute := range token.Attributes {
		mapName := attribute.TraitType + "_" + attribute.Value
		trait, ok := traitMap[mapName]
		if ok {
			withFloor = append(withFloor, trait)
		}
	}

	sort.Slice(withFloor, func(i, j int) bool {
		return withFloor[i].FloorPrice.Price > withFloor[j].FloorPrice.Price
	})

	return withFloor

}

func (a *assetService) GetTokenIds(assets []*models.AssetEntity) []string {
	var tokenIds []string

	for _, a := range assets {
		tokenIds = append(tokenIds, a.TokenID)
	}

	return tokenIds
}

func (a *assetService) HandleAssets(request dto.AssetRequest) []*models.AssetEntity {
	var task *models.TaskAssets

	in := make(chan *models.TaskSingleAsset)
	out := make(chan *models.TaskSingleAsset)

	url := os.Getenv("FIRE_PROX")

	assets := a.GetAssets(url, request.Slug, request.Owner)
	if len(assets.Assets) == 0 {
		for i := 0; i < 50; i++ {
			fmt.Println("---")
			log.Println("RETRY: ", i)
			log.Println("SLUG: ", request.Slug)
			fmt.Println("---")
			assets = a.GetAssets(url, request.Slug, request.Owner)
			if len(assets.Assets) != 0 {
				break
			}
		}
	}
	var floorPrice float64

	coll, err := a.dao.Repo.NewMetaQuery().CheckIfScraped(request.Slug)
	if err != nil {
		stats := a.GetStats(url, request.Slug)
		if stats.Stats.Count == 0 {
			for i := 0; i < 50; i++ {
				stats = a.GetStats(url, request.Slug)
				if stats.Stats.Count > 0 {
					break
				}
			}
		}
		floorPrice = stats.Stats.FloorPrice
		task = &models.TaskAssets{
			Assets:     assets.Assets,
			FloorPrice: floorPrice,
		}

		for i := 0; i < 15; i++ {
			go utils.Worker(in, out)
		}
		go func() {
			a.Producer(in, task)
		}()

		assets := a.AssetConsumer(out)
		return assets
	} else {
		floorPrice = coll.FloorPrice.Price
		traits, err := a.dao.Repo.NewTraitQuery().GetTraits(coll.Collection)
		if err != nil {
			return nil
		}
		task = &models.TaskAssets{
			Assets:     assets.Assets,
			FloorPrice: floorPrice,
			Collection: coll.Collection,
			Traits:     traits,
		}
		for i := 0; i < 15; i++ {
			go utils.Worker(in, out)
		}
		go func() {
			a.Producer(in, task)
		}()

		assets := a.ScrapedAssetsConsumer(out)

		return assets
	}
}
