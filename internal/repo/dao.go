package repo

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"time"
)

var (
	metaCollection  = "META-collections"
	traitCollection = "-all-traits"
)

type DAO interface {
	NewMetaQuery() MetaQuery
	NewCollectionQuery() CollectionQuery
	NewTraitQuery() TraitQuery
}

type dao struct {
}

var DB *mongo.Database

func NewDAO(db *mongo.Database) DAO {
	DB = db
	return &dao{}
}

func NewMongo() (*mongo.Database, error) {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			fmt.Println("Could not load .env file")
			return nil, err
		}
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		fmt.Println("Invalid uri")
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("Failed to connect to mongo")
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("Failed to ping mongo")
		return nil, err
	}

	log.Println("Connection to mongo successful")

	db := client.Database(os.Getenv("DB"))

	return db, nil

}

func (d *dao) NewMetaQuery() MetaQuery {
	return &metaQuery{}
}

func (d *dao) NewCollectionQuery() CollectionQuery {
	return &collectionQuery{}
}

func (d *dao) NewTraitQuery() TraitQuery {
	return &traitQuery{}
}
