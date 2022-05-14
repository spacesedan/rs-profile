package dto

// AssetRequest format of the request body sent from the client
type AssetRequest struct {
	Owner           string `json:"owner" form:"owner" validate:"required" binding:"required"`
	ContractAddress string `json:"contractAddress" form:"contractAddress" validate:"required" binding:"required"`
}

type OwnedCollectionsRequest struct {
	Collection []string `form:"collection" validate:"required" binding:"required"`
}

type CollectionInformationRequest struct {
	ContractAddress string `form:"contractAddress" validate:"required" binding:"required"`
}

// CollectionRequest the request used in the collections endpoint
type CollectionRequest struct {
	Owner   string `form:"owner"`
	Refresh bool   `form:"refresh"`
}

type MetadataRequest struct {
	ContractAddress string `json:"contractAddress" uri:"contractAddress" validate:"required" binding:"required"`
	Owner           string `json:"owner" uri:"owner" validate:"required" binding:"required"`
}

type AssetsWithRefresh struct {
	ContractAddress string `form:"contractAddress" validate:"required"`
	Owner           string `form:"owner" validate:"required"`
	Refresh         bool   `form:"refresh"`
}
