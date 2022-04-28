package models

type DBToken struct {
	Name            string       `bson:"name"`
	Number          int          `bson:"number"`
	RarityScoreRank int          `bson:"rarityScoreRank"`
	Attributes      []Attributes `bson:"attributes"`
}

type Attributes struct {
	TraitType string `bson:"trait_type"`
	Value     string `bson:"value"`
}
