package repo

import (
	"context"
	"github.com/spacesedan/profile-tracker/internal/deps"
	"github.com/spacesedan/profile-tracker/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type TraitQuery interface {
	GetTraits(collectionName string) (map[string]*models.DBTraits, error)
}

type traitQuery struct {
}

func (t *traitQuery) GetTraits(collectionName string) (map[string]*models.DBTraits, error) {
	m := DB.Collection(collectionName + "-all-traits")

	ctx := context.Background()

	var traits []*models.DBTraits

	results, err := m.Find(ctx, bson.M{})
	if err != nil {
		log.Fatalln(err)
	}
	if results == nil {
		return nil, deps.ErrNothingFound
	}
	defer func(results *mongo.Cursor, ctx context.Context) {
		err := results.Close(ctx)
		if err != nil {
			return
		}
	}(results, ctx)
	for results.Next(ctx) {
		var trait *models.DBTraits
		results.Decode(&trait)
		traits = append(traits, trait)
	}
	traitMap := make(map[string]*models.DBTraits)
	for i := range traits {
		mapName := traits[i].ID.TraitType + "_" + traits[i].ID.Value
		traitMap[mapName] = traits[i]
	}

	return traitMap, nil
}
