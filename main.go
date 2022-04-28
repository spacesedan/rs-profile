package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	ssm_client "github.com/spacesedan/profile-tracker/pkg/config/aws"
	"github.com/spacesedan/profile-tracker/pkg/config/mongo"
	"github.com/spacesedan/profile-tracker/pkg/server/routes"
	"log"
	"os"
)

func main() {
	ssm_client.NewSSMClient()
	mongo.ConnectDb()

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	routes.SetupRoutes(app)

	port := os.Getenv("PORT")

	err := app.Listen(":" + port)
	if err != nil {
		log.Fatalln(err)
	}
}
