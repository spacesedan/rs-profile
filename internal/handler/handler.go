package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/spacesedan/profile-tracker/internal/datastores"
	"github.com/spacesedan/profile-tracker/internal/service"
)

type Handler struct {
	AssetService      service.AssetService
	CollectionService service.CollectionService
	MetadataService   service.MetadataService
	Cache             datastores.Cache
}

type Config struct {
	Router            *gin.Engine
	AssetService      service.AssetService
	CollectionService service.CollectionService
	MetadataService   service.MetadataService
	Cache             datastores.Cache
}

func NewHandler(c Config) {
	h := &Handler{
		AssetService:      c.AssetService,
		CollectionService: c.CollectionService,
		MetadataService:   c.MetadataService,
		Cache:             c.Cache,
	}

	c.Router.GET("/health", h.Health)
	c.Router.GET("/v1/api/assets", h.GetOwned)
	c.Router.GET("/v1/api/collections", h.Collections)
	c.Router.GET("/v1/api/collections/owned", h.OwnedCollections)
	c.Router.GET("/v1/api/collections/get", h.GetCollection)
	c.Router.GET("/v1/api/metadata/:contractAddress/:walletAddress", h.Metadata)

}
