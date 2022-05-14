package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spacesedan/profile-tracker/internal/datastores"
	"github.com/spacesedan/profile-tracker/internal/handler"
	"github.com/spacesedan/profile-tracker/internal/service"
)

func inject(dao *datastores.DAO) (*gin.Engine, error) {

	asset := service.NewAssetService(dao)
	collection := service.NewCollectionService(dao)
	metadata := service.NewMetadataService(dao)
	app := gin.Default()

	//	Slug:  "little-lemon-friends",
	//	Owner: "0xd5A771Da32A392036a98f7DA6b11D46D6D1c61f9",
	//	//Owner:  "0x3a1BF0c3395975E571Cd78B9191819FF1B015A50",
	//	Cursor: "",

	app.Use(cors.New(
		cors.Config{
			AllowAllOrigins:  true,
			AllowCredentials: true,
		}))

	handler.NewHandler(handler.Config{
		Router:            app,
		AssetService:      asset,
		CollectionService: collection,
		MetadataService:   metadata,
		Cache:             dao.Cache,
	})

	return app, nil
}
