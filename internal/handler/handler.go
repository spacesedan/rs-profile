package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/spacesedan/profile-tracker/internal/service"
)

type Handler struct {
	AssetService      service.AssetService
	CollectionService service.CollectionService
	MetadataService   service.MetadataService
}

type Config struct {
	Router            *gin.Engine
	AssetService      service.AssetService
	CollectionService service.CollectionService
	MetadataService   service.MetadataService
}

func NewHandler(c Config) {
	h := &Handler{
		AssetService:      c.AssetService,
		CollectionService: c.CollectionService,
		MetadataService:   c.MetadataService,
	}

	c.Router.GET("/health", h.Health)
	c.Router.GET("/api/v1/assets", h.Assets)
	c.Router.GET("/api/v1/collections", h.Collections)
	c.Router.GET("/api/v1/assets/owned", h.GetOwned)
	c.Router.GET("/api/v1/metadata/:contractAddress/:walletAddress", h.Metadata)

}
