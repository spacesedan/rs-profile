package models

type OSCollection struct {
	Collections []Collection `json:"collections""`
}

type Collection struct {
	Ranked                bool                   `json:"ranked"`
	Ignore                bool                   `json:"ignore"`
	ContractAddress       string                 `json:"contractAddress"'`
	PrimaryAssetContracts []PrimaryAssetContract `json:"primary_asset_contracts"`
	Traits                struct {
	} `json:"traits"`
	Stats          OSStats `json:"stats"`
	BannerImageURL string  `json:"banner_image_url"`
	ChatURL        interface {
	} `json:"chat_url"`
	CreatedDate             string `json:"created_date"`
	DefaultToFiat           bool   `json:"default_to_fiat"`
	Description             string `json:"description"`
	DevBuyerFeeBasisPoints  string `json:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints string `json:"dev_seller_fee_basis_points"`
	DiscordURL              string `json:"discord_url"`
	DisplayData             struct {
		CardDisplayStyle string `json:"card_display_style"`
	} `json:"display_data"`
	ExternalURL           string `json:"external_url"`
	Featured              bool   `json:"featured"`
	FeaturedImageURL      string `json:"featured_image_url"`
	Hidden                bool   `json:"hidden"`
	SafelistRequestStatus string `json:"safelist_request_status"`
	ImageURL              string `json:"image_url"`
	IsSubjectToWhitelist  bool   `json:"is_subject_to_whitelist"`
	LargeImageURL         string `json:"large_image_url"`
	MediumUsername        interface {
	} `json:"medium_username"`
	Name                        string `json:"name"`
	OnlyProxiedTransfers        bool   `json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  string `json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints string `json:"opensea_seller_fee_basis_points"`
	PayoutAddress               string `json:"payout_address"`
	RequireEmail                bool   `json:"require_email"`
	ShortDescription            interface {
	} `json:"short_description"`
	Slug        string `json:"slug"`
	TelegramURL interface {
	} `json:"telegram_url"`
	TwitterUsername interface {
	} `json:"twitter_username"`
	InstagramUsername string `json:"instagram_username"`
	WikiURL           interface {
	} `json:"wiki_url"`
	IsNsfw          bool `json:"is_nsfw"`
	OwnedAssetCount int  `json:"owned_asset_count"`
}

type PrimaryAssetContract struct {
	Address                     string      `json:"address"`
	AssetContractType           string      `json:"asset_contract_type"`
	CreatedDate                 string      `json:"created_date"`
	Name                        string      `json:"name"`
	NftVersion                  string      `json:"nft_version"`
	OpenseaVersion              interface{} `json:"opensea_version"`
	Owner                       int         `json:"owner"`
	SchemaName                  string      `json:"schema_name"`
	Symbol                      string      `json:"symbol"`
	TotalSupply                 string      `json:"total_supply"`
	Description                 string      `json:"description"`
	ExternalLink                string      `json:"external_link"`
	ImageURL                    string      `json:"image_url"`
	DefaultToFiat               bool        `json:"default_to_fiat"`
	DevBuyerFeeBasisPoints      int         `json:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints     int         `json:"dev_seller_fee_basis_points"`
	OnlyProxiedTransfers        bool        `json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  int         `json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints int         `json:"opensea_seller_fee_basis_points"`
	BuyerFeeBasisPoints         int         `json:"buyer_fee_basis_points"`
	SellerFeeBasisPoints        int         `json:"seller_fee_basis_points"`
	PayoutAddress               string      `json:"payout_address"`
}
