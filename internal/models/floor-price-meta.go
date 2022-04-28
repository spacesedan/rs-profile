package models

type Meta struct {
	FloorPrice FloorPriceMeta `bson:"floorPrice" json:"floorPrice"`
	Collection string         `bson:"collection" json:"collection"`
	Slug       string         `bson:"slug" json:"slug"`
}

type FloorPriceMeta struct {
	Price          float64 `bson:"price" json:"price"`
	PriceEntryTime string  `bson:"priceEntryTime" json:"priceEntryTime"`
}
