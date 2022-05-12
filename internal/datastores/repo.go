package datastores

import (
	"context"
	"fmt"
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

type Repo interface {
	NewMetaQuery() MetaQuery
	NewCollectionQuery() CollectionQuery
	NewTraitQuery() TraitQuery
}

type repo struct {
}

var DB *mongo.Database

func NewRepo(db *mongo.Database) Repo {
	DB = db
	return &repo{}
}

func NewMongo() (*mongo.Database, error) {

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

func (r *repo) NewMetaQuery() MetaQuery {
	return &metaQuery{}
}

func (r *repo) NewCollectionQuery() CollectionQuery {
	return &collectionQuery{}
}

func (r *repo) NewTraitQuery() TraitQuery {
	return &traitQuery{}
}
