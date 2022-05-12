package models

type OwnedMeta struct {
	CollectionName string  `json:"collectionName" bson:"collection"`
	DisplayName    string  `json:"displayName" bson:"displayName"`
	FloorPrice     float64 `json:"floorPrice" bson:"price"`
	BannerImage    string  `json:"bannerImage" bson:"bannerUrl"`
	Slug           string  `json:"slug" bson:"slug"`
	ImageUrl       string  `json:"imageUrl" bson:"imageUrl"`
}
