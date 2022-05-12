package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Meta struct {
	ID                     primitive.ObjectID `json:"_id" bson:"_id"`
	Collection             string             `json:"collection" bson:"collection"`
	Contract               string             `json:"contract" bson:"contract"`
	DisplayName            string             `json:"displayName" bson:"displayName"`
	NoneValue              interface{}        `json:"noneValue" bson:"noneValue"`
	ContractType           string             `json:"contractType" bson:"contractType"`
	Slug                   string             `json:"slug" bson:"slug"`
	CreatedDate            primitive.DateTime `json:"created_date" bson:"created_date"`
	BannerURL              string             `json:"bannerUrl" bson:"bannerUrl"`
	ImageURL               string             `json:"imageUrl" bson:"imageUrl"`
	Description            string             `json:"description" bson:"description"`
	Volume                 int                `json:"volume" bson:"volume"`
	Owners                 int                `json:"owners" bson:"owners"`
	InProgress             bool               `json:"inProgress" json:"inProgress"`
	PriceScraping          bool               `json:"priceScraping" bson:"priceScraping"`
	FloorPrice             FloorPrice         `json:"floorPrice" bson:"floorPrice"`
	AttributeTypeBreakdown []struct {
		Type  string `json:"type" bson:"type"`
		Count int    `json:"count" bson:"count"`
	} `json:"attributeTypeBreakdown" bson:"attributeTypeBreakdown"`
	AttributeBreakdown []struct {
		TraitCount int `json:"traitCount" bson:"traitCount"`
		Occurrence int `json:"occurrence" bson:"occurrence"`
	} `json:"attributeBreakdown" bson:"attributeBreakdown"`
	AttributeTypeOccuranceBreakdown []struct {
		Type  string `json:"type" bson:"type"`
		Count int    `json:"count" bson:"count"`
	} `json:"attributeTypeOccuranceBreakdown" bson:"attributeTypeOccuranceBreakdown"`
	StartingBlock int `json:"startingBlock" bson:"startingBlock"`
}

type FloorPrice struct {
	Price          float64            `json:"price" bson:"price"`
	PriceEntryTime primitive.DateTime `json:"priceEntryTime" bson:"priceEntryTime"`
}
