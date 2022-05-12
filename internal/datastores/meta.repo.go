package datastores

import (
	"context"
	"fmt"
	"github.com/spacesedan/profile-tracker/internal/deps"
	"github.com/spacesedan/profile-tracker/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MetaQuery interface {
	CheckIfScraped(slug string) (*models.Meta, error)
	GetAllScraped() (map[string]*models.Meta, error)
	GetByName(req []string) []*models.OwnedMeta
}

type metaQuery struct {
}

func (m *metaQuery) CheckIfScraped(slug string) (*models.Meta, error) {
	c := DB.Collection(metaCollection)

	var coll *models.Meta

	ctx := context.Background()

	result := c.FindOne(ctx, bson.D{{"slug", slug}})
	if result == nil {
		fmt.Println("nothing found")
		return nil, deps.ErrNothingFound
	}
	err := result.Decode(&coll)
	if err != nil {
		return nil, deps.ErrNothingFound
	}

	return coll, nil

}

func (m *metaQuery) GetByName(req []string) []*models.OwnedMeta {
	c := DB.Collection("META-collections")
	ctx := context.Background()

	sort := bson.D{
		{"$sort", bson.M{
			"floorPrice.price": -1,
		}},
	}

	match := bson.D{
		{"$match", m.buildOwnedQuery(req)},
	}

	project := bson.D{
		{"$project", bson.M{
			"bannerUrl":        true,
			"imageUrl":         true,
			"slug":             true,
			"floorPrice.price": true,
			"collection":       true,
			"displayName":      true,
		}},
	}

	opts := options.Aggregate()
	opts.SetAllowDiskUse(true)

	var collections []*models.Meta

	aggregate, err := c.Aggregate(ctx, mongo.Pipeline{sort, match, project}, opts)
	if err != nil {
		log.Printf("Error while querying owned collections: %v", err)
		return nil
	}

	if aggregate.Err() != nil {
		log.Printf("Unknown internal error: %v", aggregate.Err())
		return nil
	}

	defer aggregate.Close(ctx)
	for aggregate.Next(ctx) {
		var collection *models.Meta
		aggregate.Decode(&collection)
		collections = append(collections, collection)
	}

	var cs []*models.OwnedMeta

	for _, c := range collections {
		cs = append(cs, &models.OwnedMeta{
			CollectionName: c.Collection,
			DisplayName:    c.DisplayName,
			FloorPrice:     c.FloorPrice.Price,
			BannerImage:    c.BannerURL,
			Slug:           c.Slug,
			ImageUrl:       c.ImageURL,
		})
	}

	return cs

}

func (m *metaQuery) GetAllScraped() (map[string]*models.Meta, error) {
	c := DB.Collection("META-collections")
	ctx := context.Background()

	var colls []*models.Meta

	all, _ := c.Find(ctx, bson.D{{}})
	defer all.Close(ctx)
	for all.Next(ctx) {
		var coll *models.Meta
		all.Decode(&coll)
		colls = append(colls, coll)
	}

	collsMap := make(map[string]*models.Meta, len(colls))
	for _, coll := range colls {
		collsMap[coll.Slug] = coll
	}

	return collsMap, nil
}

func (m *metaQuery) buildOwnedQuery(collections []string) bson.M {
	var queryArr []bson.M
	for _, collection := range collections {
		query := bson.M{"slug": collection}
		queryArr = append(queryArr, query)
	}

	return bson.M{"$or": queryArr}

}
