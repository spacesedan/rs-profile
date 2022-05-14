package service

import (
	"github.com/spacesedan/profile-tracker/internal/datastores"
	"github.com/spacesedan/profile-tracker/internal/dto"
	"github.com/spacesedan/profile-tracker/internal/models"
	"github.com/spacesedan/profile-tracker/internal/utils"
	"log"
	"sort"
	"strconv"
	"sync"
	"time"
)

type AssetService interface {
	GetTokenCount(req dto.AssetsWithRefresh) int
	GetFloorPrice(req dto.AssetsWithRefresh) float64
	GetTokens(req dto.AssetsWithRefresh, count int) []models.ReservoirToken
	GetOwnedTokens(req dto.AssetsWithRefresh) dto.GetOwnedResponse
}

type assetService struct {
	dao *datastores.DAO
}

func NewAssetService(dao *datastores.DAO) AssetService {
	return &assetService{
		dao: dao,
	}
}

func (a *assetService) GetTokenCount(req dto.AssetsWithRefresh) int {
	var assets models.MoralisAssetResponse
	uri := "https://deep-index.moralis.io/api/v2/" + req.Owner + "/nft/" + req.ContractAddress + "?chain=eth&format=decimal"
	utils.MakeMoralisGetRequest(uri, &assets)
	return assets.Total
}

func (a *assetService) GetFloorPrice(req dto.AssetsWithRefresh) float64 {
	var priceMap models.ReservoirPriceMap
	url := "https://api.reservoir.tools/tokens/floor/v1?contract=" + req.ContractAddress
	utils.GetJson(url, &priceMap)

	var floorPrices []float64

	for _, element := range priceMap.Tokens {
		if element != 0 {
			floorPrices = append(floorPrices, element)
		}
	}

	sort.Slice(floorPrices, func(a, b int) bool {
		return floorPrices[a] < floorPrices[b]
	})

	return floorPrices[0]

}

func (a *assetService) GetTokens(req dto.AssetsWithRefresh, count int) []models.ReservoirToken {
	var tokens []models.ReservoirToken

	pages := count / 20

	if pages == 0 {
		var offset int
		var reservoirAssets models.ReservoirTokensResponse
		url := "https://api.reservoir.tools/users/" + req.Owner + "/tokens/v2?collection=" + req.ContractAddress + "&offset=" + strconv.Itoa(offset) + "&limit=20"
		log.Printf(url)
		utils.GetJson(url, &reservoirAssets)
		for _, t := range reservoirAssets.Tokens {
			tokens = append(tokens, t.Token)
		}
		return tokens
	}

	in := make(chan int, pages)
	out := make(chan models.ReservoirTokensResponse, pages)

	for w := 0; w <= pages; w++ {
		go func(w int) {
			for i := range in {
				var reservoirAssets models.ReservoirTokensResponse
				offset := i * 20
				url := "https://api.reservoir.tools/users/" + req.Owner + "/tokens/v2?collection=" + req.ContractAddress + "&offset=" + strconv.Itoa(offset) + "&limit=20"
				utils.GetJson(url, &reservoirAssets)
				out <- reservoirAssets
			}
		}(w)
	}

	for j := 0; j <= pages; j++ {
		go func(j int) {
			in <- j
		}(j)
	}

	for {
		select {
		case msg := <-out:
			for _, token := range msg.Tokens {
				tokens = append(tokens, token.Token)
			}
		case <-time.After(300 * time.Millisecond):
			return tokens
		}
	}

}

func (a *assetService) GetOwnedTokens(req dto.AssetsWithRefresh) dto.GetOwnedResponse {
	var wg sync.WaitGroup
	var tokenCount int
	var floorPrice float64

	wg.Add(2)

	go func() {
		defer wg.Done()
		floorPrice = a.GetFloorPrice(req)
	}()

	go func() {
		defer wg.Done()
		tokenCount = a.GetTokenCount(req)
	}()
	wg.Wait()

	tokens := a.GetTokens(req, tokenCount)

	return dto.GetOwnedResponse{
		Tokens:     tokens,
		FloorPrice: floorPrice,
	}

}

//func (a *assetService) GetStats(url, slug string) models.OSStats {
//	uri := url + "collection/" + slug + "/stats"
//	var stats models.OSStats
//	utils.GetJson(uri, &stats)
//
//	return stats
//}
//
//func (a *assetService) GetRarestTrait(asset *models.AssetEntity) *models.DBTraits {
//	sort.Slice(asset.Traits, func(i, j int) bool {
//		return asset.Traits[i].TraitCount > asset.Traits[j].TraitCount
//	})
//
//	var rarest = &models.DBTraits{}
//
//	rarest.ID.TraitType = asset.Traits[0].TraitType
//	rarest.ID.Value = asset.Traits[0].Value
//
//	return rarest
//}
//
//func (a *assetService) Producer(ch chan<- *models.TaskSingleAsset, task *models.TaskAssets) {
//	for _, asset := range task.Assets {
//		var single *models.TaskSingleAsset
//		single = &models.TaskSingleAsset{
//			Asset:      asset,
//			Collection: task.Collection,
//			FloorPrice: task.FloorPrice,
//			Traits:     task.Traits,
//		}
//		ch <- single
//	}
//}
//
//func (a *assetService) AssetConsumer(chA <-chan *models.TaskSingleAsset) []*models.AssetEntity {
//	var owned []*models.AssetEntity
//
//	for {
//		select {
//		case msg := <-chA:
//			if len(msg.Asset.Traits) != 0 {
//				rarestTrait := a.GetRarestTrait(msg.Asset)
//				msg.Asset.TopTrait = rarestTrait
//			}
//			fmt.Printf("%+v\n", msg.Asset.TopTrait)
//			msg.Asset.FloorPrice = msg.FloorPrice
//			msg.Asset.TraitFloorPrice = msg.FloorPrice
//			msg.Asset.TopTrait.RarityScore = 0
//			msg.Asset.TopTrait.FloorPrice.Price = 0
//			msg.Asset.Rank = 0
//			msg.Asset.TopTrait.FloorPrice.PriceEntryTime = primitive.NewDateTimeFromTime(time.Now())
//			msg.Asset.DBName = "null"
//			owned = append(owned, msg.Asset)
//		case <-time.After(50 * time.Millisecond):
//			return owned
//		}
//	}
//}
//
//func (a *assetService) ScrapedAssetsConsumer(chA <-chan *models.TaskSingleAsset) []*models.AssetEntity {
//	var owned []*models.AssetEntity
//
//	for {
//		select {
//		case msg := <-chA:
//			token := a.dao.Repo.NewCollectionQuery().GetToken(msg.Collection, msg.Asset.TokenID)
//			traitsWithValue := a.TraitsWithFloor(msg.Traits, token)
//			msg.Asset.FloorPrice = msg.FloorPrice
//			msg.Asset.TraitFloorPrice = traitsWithValue[0].FloorPrice.Price
//			msg.Asset.TopTrait = traitsWithValue[0]
//			msg.Asset.DBName = msg.Collection
//			msg.Asset.Rank = token.RarityScoreRank
//			owned = append(owned, msg.Asset)
//		case <-time.After(50 * time.Millisecond):
//			return owned
//		}
//	}
//}
//
//func (a *assetService) TraitsWithFloor(traitMap map[string]*models.DBTraits, token *models.DBToken) []*models.DBTraits {
//	var withFloor []*models.DBTraits
//
//	for _, attribute := range token.Attributes {
//		mapName := attribute.TraitType + "_" + attribute.Value
//		trait, ok := traitMap[mapName]
//		if ok {
//			withFloor = append(withFloor, trait)
//		}
//	}
//
//	sort.Slice(withFloor, func(i, j int) bool {
//		return withFloor[i].FloorPrice.Price > withFloor[j].FloorPrice.Price
//	})
//
//	return withFloor
//
//}
//
//func (a *assetService) HandleAssets(request dto.AssetRequest) []*models.AssetEntity {
//	var task *models.TaskAssets
//
//	in := make(chan *models.TaskSingleAsset)
//	out := make(chan *models.TaskSingleAsset)
//
//	assets := a.GetAssets(request.Slug, request.Owner)
//	if len(assets.Assets) == 0 {
//		for i := 0; i < 50; i++ {
//			fmt.Println("---")
//			log.Println("RETRY: ", i)
//			log.Println("SLUG: ", request.Slug)
//			fmt.Println("---")
//			assets = a.GetAssets(request.Slug, request.Owner)
//			if len(assets.Assets) != 0 {
//				break
//			}
//		}
//	}
//	var floorPrice float64
//
//	coll, err := a.dao.Repo.NewMetaQuery().CheckIfScraped(request.Slug)
//	if err != nil {
//		stats := a.GetStats(url, request.Slug)
//		if stats.Stats.Count == 0 {
//			for i := 0; i < 50; i++ {
//				stats = a.GetStats(url, request.Slug)
//				if stats.Stats.Count > 0 {
//					break
//				}
//			}
//		}
//		floorPrice = stats.Stats.FloorPrice
//		task = &models.TaskAssets{
//			Assets:     assets.Assets,
//			FloorPrice: floorPrice,
//		}
//
//		for i := 0; i < 15; i++ {
//			go utils.Worker(in, out)
//		}
//		go func() {
//			a.Producer(in, task)
//		}()
//
//		assets := a.AssetConsumer(out)
//		return assets
//	} else {
//		floorPrice = coll.FloorPrice.Price
//		traits, err := a.dao.Repo.NewTraitQuery().GetTraits(coll.Collection)
//		if err != nil {
//			return nil
//		}
//		task = &models.TaskAssets{
//			Assets:     assets.Assets,
//			FloorPrice: floorPrice,
//			Collection: coll.Collection,
//			Traits:     traits,
//		}
//		for i := 0; i < 15; i++ {
//			go utils.Worker(in, out)
//		}
//		go func() {
//			a.Producer(in, task)
//		}()
//
//		assets := a.ScrapedAssetsConsumer(out)
//
//		return assets
//	}
//}
