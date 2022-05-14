package models

type ReservoirPriceMap struct {
	Tokens map[string]float64
}

type ReservoirTokensResponse struct {
	Tokens []struct {
		Token     ReservoirToken `json:"token"`
		Ownership struct {
			TokenCount    string      `json:"tokenCount"`
			OnSaleCount   string      `json:"onSaleCount"`
			FloorAskPrice interface{} `json:"floorAskPrice"`
		} `json:"ownership"`
	} `json:"tokens"`
}

type ReservoirToken struct {
	Contract   string `json:"contract"`
	TokenID    string `json:"tokenId"`
	Name       string `json:"name"`
	Image      string `json:"image"`
	Collection struct {
		ID            string  `json:"id"`
		Name          string  `json:"name"`
		ImageURL      string  `json:"imageUrl"`
		FloorAskPrice float64 `json:"floorAskPrice"`
	} `json:"collection"`
	TopBid struct {
		ID    interface{} `json:"id"`
		Value interface{} `json:"value"`
	} `json:"topBid"`
}

type ReservoirCollectionResponse struct {
	Collection struct {
		ID       string `json:"id"`
		Slug     string `json:"slug"`
		Name     string `json:"name"`
		Metadata struct {
			ImageURL        string      `json:"imageUrl"`
			DiscordURL      string      `json:"discordUrl"`
			Description     string      `json:"description"`
			ExternalURL     string      `json:"externalUrl"`
			BannerImageURL  string      `json:"bannerImageUrl"`
			TwitterUsername interface{} `json:"twitterUsername"`
		} `json:"metadata"`
		SampleImages    []interface{} `json:"sampleImages"`
		TokenCount      string        `json:"tokenCount"`
		OnSaleCount     string        `json:"onSaleCount"`
		PrimaryContract string        `json:"primaryContract"`
		TokenSetID      string        `json:"tokenSetId"`
		Royalties       struct {
			Bps       int    `json:"bps"`
			Recipient string `json:"recipient"`
		} `json:"royalties"`
		LastBuy struct {
			Value     interface{} `json:"value"`
			Timestamp interface{} `json:"timestamp"`
		} `json:"lastBuy"`
		FloorAsk struct {
			ID         string  `json:"id"`
			Price      float64 `json:"price"`
			Maker      string  `json:"maker"`
			ValidFrom  int     `json:"validFrom"`
			ValidUntil int     `json:"validUntil"`
			Token      struct {
				Contract string      `json:"contract"`
				TokenID  string      `json:"tokenId"`
				Name     interface{} `json:"name"`
				Image    interface{} `json:"image"`
			} `json:"token"`
		} `json:"floorAsk"`
		TopBid struct {
			ID         interface{} `json:"id"`
			Value      interface{} `json:"value"`
			Maker      interface{} `json:"maker"`
			ValidFrom  interface{} `json:"validFrom"`
			ValidUntil interface{} `json:"validUntil"`
		} `json:"topBid"`
		Rank struct {
			OneDay    interface{} `json:"1day"`
			SevenDay  int         `json:"7day"`
			Three0Day int         `json:"30day"`
			AllTime   int         `json:"allTime"`
		} `json:"rank"`
		Volume struct {
			OneDay    int     `json:"1day"`
			SevenDay  float64 `json:"7day"`
			Three0Day float64 `json:"30day"`
			AllTime   float64 `json:"allTime"`
		} `json:"volume"`
		VolumeChange struct {
			OneDay    int         `json:"1day"`
			SevenDay  float64     `json:"7day"`
			Three0Day interface{} `json:"30day"`
		} `json:"volumeChange"`
		FloorSale struct {
			OneDay    float64     `json:"1day"`
			SevenDay  float64     `json:"7day"`
			Three0Day interface{} `json:"30day"`
		} `json:"floorSale"`
		FloorSaleChange struct {
			OneDay int `json:"1day"`
		} `json:"floorSaleChange"`
		CollectionBidSupported bool          `json:"collectionBidSupported"`
		OwnerCount             int           `json:"ownerCount"`
		Attributes             []interface{} `json:"attributes"`
	} `json:"collection"`
}
