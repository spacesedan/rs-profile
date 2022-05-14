package dto

import "github.com/spacesedan/profile-tracker/internal/models"

type GetOwnedResponse struct {
	Tokens     []models.ReservoirToken `json:"tokens"`
	FloorPrice float64                 `json:"floorPrice"`
}

type GetCollectionInformationResponse []CollectionInfo

type CollectionInfo struct {
	FloorPrice      float64 `json:"floorPrice"`
	Image           string  `json:"image"`
	Banner          string  `json:"banner"`
	Slug            string  `json:"slug"`
	Name            string  `json:"name"`
	ContractAddress string  `json:"contractAddress"`
}
