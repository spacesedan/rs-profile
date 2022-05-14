package models

import "time"

type MoralisAssetResponse struct {
	Total    int    `json:"total"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	Cursor   string `json:"cursor"`
	Result   []struct {
		TokenAddress      string    `json:"token_address"`
		TokenID           string    `json:"token_id"`
		OwnerOf           string    `json:"owner_of"`
		BlockNumber       string    `json:"block_number"`
		BlockNumberMinted string    `json:"block_number_minted"`
		TokenHash         string    `json:"token_hash"`
		Amount            string    `json:"amount"`
		ContractType      string    `json:"contract_type"`
		Name              string    `json:"name"`
		Symbol            string    `json:"symbol"`
		TokenURI          string    `json:"token_uri"`
		Metadata          string    `json:"metadata"`
		SyncedAt          time.Time `json:"synced_at"`
	} `json:"result"`
	Status string `json:"status"`
}
