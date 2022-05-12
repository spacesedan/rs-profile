package service

import (
	"github.com/spacesedan/profile-tracker/internal/datastores"
	"github.com/spacesedan/profile-tracker/internal/dto"
	"github.com/spacesedan/profile-tracker/internal/models"
	"github.com/spacesedan/profile-tracker/internal/utils"
	"log"
	"net/url"
)

type CollectionService interface {
	GetCollections(req dto.CollectionRequest) []string
	GetContractAddress(url, slug string) string
	GetOwnedCollection(values url.Values) []*models.OwnedMeta
}

type collectionService struct {
	dao *datastores.DAO
}

func NewCollectionService(dao *datastores.DAO) CollectionService {
	return &collectionService{
		dao: dao,
	}
}

func (c *collectionService) GetCollections(req dto.CollectionRequest) []string {
	url := "https://api.opensea.io/api/v1/collections?asset_owner=" + req.Owner + "&offset=0&limit=300"
	var collections []models.Collection
	utils.MakeOpenSeaRequest(url, &collections)

	var collectionSlugs []string

	for _, col := range collections {
		collectionSlugs = append(collectionSlugs, col.Slug)
	}

	return collectionSlugs
}

func (c *collectionService) GetOwnedCollection(values url.Values) []*models.OwnedMeta {
	collections := values["collection"]

	log.Printf("COLLECTIONS: %v", collections)
	return c.dao.Repo.NewMetaQuery().GetByName(collections)
}

func (c *collectionService) GetContractAddress(url, slug string) string {
	var asset models.Assets
	uri := url + "assets?collection=" + slug + "&order_by=pk&order_direction=desc&limit=1"
	utils.GetJson(uri, &asset)
	return asset.Assets[0].AssetContract.Address
}
