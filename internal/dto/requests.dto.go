package dto

// AssetRequest format of the request body sent from the client
type AssetRequest struct {
	Owner string `json:"owner" form:"owner" validate:"required" binding:"required"`
	Slug  string `json:"slug" form:"slug" validate:"required" binding:"required"`
}

// CollectionRequest the request used in the collections endpoint
type CollectionRequest struct {
	Owner  string `form:"wallet_address"`
	Cursor string `json:"cursor"`
}

type MetadataRequest struct {
	ContractAddress string `json:"contractAddress" uri:"contractAddress" validate:"required" binding:"required"`
	Owner           string `json:"owner" uri:"owner" validate:"required" binding:"required"`
}

type AssetsWithCursorRequest struct {
	Slug   string `form:"slug" validate:"required"`
	Owner  string `form:"owner" validate:"required"`
	Cursor string
}
