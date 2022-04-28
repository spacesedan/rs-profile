package models

type DBTraits struct {
	ID          ID             `bson:"_id" json:"id"`
	FloorPrice  FloorPriceMeta `bson:"floorPrice" json:"floorPrice"`
	RarityScore float64        `bson:"rarityScore" json:"rarityScore"`
}

type ID struct {
	TraitType string `bson:"trait_type" json:"traitType"`
	Value     string `bson:"value" json:"value"`
}
