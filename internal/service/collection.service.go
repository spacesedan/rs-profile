package service

import (
	"github.com/spacesedan/profile-tracker/internal/datastores"
	"github.com/spacesedan/profile-tracker/internal/dto"
	"github.com/spacesedan/profile-tracker/internal/models"
	"github.com/spacesedan/profile-tracker/internal/utils"
	"log"
	"net/url"
	"time"
)

type CollectionService interface {
	GetCollections(req dto.CollectionRequest) []dto.CollectionInfo
	GetCollectionContractAddresses(req dto.CollectionRequest) []string
	GetContractAddress(url, slug string) string
	GetOwnedCollection(values url.Values) []*models.OwnedMeta
	GetCollectionInformation(req dto.CollectionInformationRequest) dto.CollectionInfo
}

const (
	OpenSeaSymbol                 = "OPENSTORE"
	ReservoirCollectionRequestURL = "https://api.reservoir.tools/collection/v2?id="
)

type collectionService struct {
	dao *datastores.DAO
}

func NewCollectionService(dao *datastores.DAO) CollectionService {
	return &collectionService{
		dao: dao,
	}
}

// GetCollectionContractAddresses get the contracts addresses for collections a user has stakes in.
func (c *collectionService) GetCollectionContractAddresses(req dto.CollectionRequest) []string {
	url := "https://api.opensea.io/api/v1/collections?asset_owner=" + req.Owner + "&offset=0&limit=300"
	var collections []models.Collection
	utils.MakeOpenSeaRequest(url, &collections)

	var collectionContracts []string

	for _, col := range collections {
		if len(col.PrimaryAssetContracts) > 0 {
			if col.PrimaryAssetContracts[0].Symbol != OpenSeaSymbol {
				collectionContracts = append(collectionContracts, col.PrimaryAssetContracts[0].Address)
			}
		}
	}

	return collectionContracts
}

func (c *collectionService) GetCollectionInformation(req dto.CollectionInformationRequest) dto.CollectionInfo {
	var reservoirCollection models.ReservoirCollectionResponse

	url := ReservoirCollectionRequestURL + req.ContractAddress
	utils.GetJson(url, &reservoirCollection)

	log.Printf("COLLECTION NAME: %v\n", reservoirCollection.Collection.Name)

	return dto.CollectionInfo{
		FloorPrice:      reservoirCollection.Collection.FloorAsk.Price,
		Image:           reservoirCollection.Collection.Metadata.ImageURL,
		Banner:          reservoirCollection.Collection.Metadata.BannerImageURL,
		Slug:            reservoirCollection.Collection.Slug,
		Name:            reservoirCollection.Collection.Name,
		ContractAddress: reservoirCollection.Collection.ID,
	}

}

func (c *collectionService) GetOwnedCollection(values url.Values) []*models.OwnedMeta {
	collections := values["collection"]

	log.Printf("COLLECTIONS: %v", collections)
	return c.dao.Repo.NewMetaQuery().GetByName(collections)
}

func (c *collectionService) GetCollections(req dto.CollectionRequest) []dto.CollectionInfo {
	var bucket []dto.CollectionInfo
	contractAddresses := c.GetCollectionContractAddresses(req)

	in := make(chan string, len(contractAddresses))
	out := make(chan dto.CollectionInfo, len(contractAddresses))

	for w := 0; w < len(contractAddresses); w++ {
		go func(w int) {
			for contractAddress := range in {
				collectionInfo := c.GetCollectionInformation(dto.CollectionInformationRequest{ContractAddress: contractAddress})

				out <- collectionInfo
			}

		}(w)
	}

	for j := 0; j < len(contractAddresses); j++ {
		go func(j int) {
			in <- contractAddresses[j]
		}(j)
	}

	for {
		select {
		case msg := <-out:
			if msg.Name != "" {
				bucket = append(bucket, msg)
			}
		case <-time.After(1 * time.Second):
			return bucket
		}
	}
}

func (c *collectionService) GetContractAddress(url, slug string) string {
	var asset models.Assets
	uri := url + "assets?collection=" + slug + "&order_by=pk&order_direction=desc&limit=1"
	utils.GetJson(uri, &asset)
	return asset.Assets[0].AssetContract.Address
}
