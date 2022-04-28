package repo

import (
	"context"
	"fmt"
	"github.com/spacesedan/profile-tracker/internal/deps"
	"github.com/spacesedan/profile-tracker/internal/models"
	"go.mongodb.org/mongo-driver/bson"
)

type MetaQuery interface {
	CheckIfScraped(slug string) (*models.Meta, error)
	GetAllScraped() (map[string]*models.Meta, error)
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
