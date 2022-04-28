package models

type TaskAssets struct {
	Assets     []AssetEntity
	Collection string
	FloorPrice float64
	Traits     map[string]*DBTraits
}

type TaskSingleAsset struct {
	Asset      AssetEntity
	Collection string
	FloorPrice float64
	Traits     map[string]*DBTraits
}

type TaskCollections struct {
	Collections []Collection
	CollMap     map[string]*Meta
}

type TaskSingleCollection struct {
	Collection Collection
	CollMap    map[string]*Meta
}
