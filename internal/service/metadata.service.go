package service

import (
	"github.com/spacesedan/profile-tracker/internal/dto"
	"github.com/spacesedan/profile-tracker/internal/models"
	"github.com/spacesedan/profile-tracker/internal/repo"
	"github.com/spacesedan/profile-tracker/internal/utils"
	"os"
	"strconv"
	"time"
)

type MetadataService interface {
	GetAssets(url, slug, walletAddress string) dto.Refresh
	HandleMetadata(request dto.MetadataRequest) string
	RefreshMetadata(tokenId, contractAddress string)
}

type metadataService struct {
	dao repo.DAO
}

func NewMetadataService(dao repo.DAO) MetadataService {
	return &metadataService{
		dao: dao,
	}
}

func (m *metadataService) GetAssets(url, contractAddress, walletAddress string) dto.Refresh {
	var assets *models.Assets
	uri := url + "assets?owner=" + walletAddress + "&asset_contract_address=" + contractAddress + "&order_by=pk&order_direction=desc&limit=50"
	utils.GetJson(uri, &assets)

	var tokenIds dto.Refresh
	for _, asset := range assets.Assets {
		tokenIds.TokenIds = append(tokenIds.TokenIds, asset.TokenID)
	}

	return tokenIds
}

func (m *metadataService) RefreshMetadata(tokenId, contractAddress string) {
	var asset *models.AssetEntity

	uri := "https://api.opensea.io/api/v1/asset/" + contractAddress + "/" + tokenId + "?force_update=true"

	utils.GetJson(uri, &asset)

}

func (m *metadataService) HandleMetadata(request dto.MetadataRequest) string {
	url := os.Getenv("FIRE_PROX")

	refresh := m.GetAssets(url, request.ContractAddress, request.Owner)

	refresh.ContractAddress = request.ContractAddress

	for _, token := range refresh.TokenIds {
		m.RefreshMetadata(token, refresh.ContractAddress)
		time.Sleep(200 * time.Millisecond)
	}

	var msg string

	length := strconv.Itoa(len(refresh.TokenIds))

	switch len(refresh.TokenIds) {
	case 0:
		msg = "Nothing found"
	case 1:
		msg = "Successfully updated metadata of " + length + " item"
	default:
		msg = "Successfully updated metadata of " + length + " items"

	}

	return msg

}
