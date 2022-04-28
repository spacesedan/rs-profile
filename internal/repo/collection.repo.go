package repo

import (
	"context"
	"github.com/spacesedan/profile-tracker/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
)

type CollectionQuery interface {
	GetToken(collectionName string, tokenId string) *models.DBToken
}

type collectionQuery struct {
}

func (c *collectionQuery) GetToken(collectionName string, tokenId string) *models.DBToken {
	tId, _ := strconv.Atoi(tokenId)
	m := DB.Collection(collectionName)
	ctx := context.Background()

	var token *models.DBToken
	result := m.FindOne(ctx, bson.M{"number": tId})
	if result == nil {
		return nil
	}

	result.Decode(&token)

	return token
}
