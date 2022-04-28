package service

import (
	"fmt"
	"github.com/spacesedan/profile-tracker/internal/dto"
	"github.com/spacesedan/profile-tracker/internal/models"
	"github.com/spacesedan/profile-tracker/internal/repo"
	"github.com/spacesedan/profile-tracker/internal/utils"
	"os"
	"time"
)

type CollectionService interface {
	GetCollections(walletAddress string) []models.Collection
	GetContractAddress(url, slug string) string
	CollectionConsumer(ch <-chan *models.TaskSingleCollection) []*models.Collection
	HandleCollections(request dto.CollectionRequest) []*models.Collection
}

type collectionService struct {
	dao repo.DAO
}

func NewCollectionService(dao repo.DAO) CollectionService {
	return &collectionService{
		dao: dao,
	}
}

func (c *collectionService) GetCollections(walletAddress string) []models.Collection {
	url := os.Getenv("FIRE_PROX") + "collections?asset_owner=" + walletAddress + "&limit=300"
	var collections []models.Collection
	utils.GetJson(url, &collections)

	return collections
}

func (c *collectionService) GetContractAddress(url, slug string) string {
	var asset models.Assets
	uri := url + "assets?collection=" + slug + "&order_by=pk&order_direction=desc&limit=1"
	utils.GetJson(uri, &asset)
	return asset.Assets[0].AssetContract.Address
}
func (c *collectionService) Producer(ch chan<- *models.TaskSingleCollection, task *models.TaskCollections) {
	for _, collection := range task.Collections {
		single := &models.TaskSingleCollection{
			Collection: collection,
			CollMap:    task.CollMap,
		}
		ch <- single
	}
}

func (c *collectionService) CollectionConsumer(ch <-chan *models.TaskSingleCollection) []*models.Collection {
	var owned []*models.Collection

	for {
		select {
		case msg := <-ch:
			if len(msg.Collection.PrimaryAssetContracts) == 0 {
				contractAddress := c.GetContractAddress(os.Getenv("FIRE_PROX"), msg.Collection.Slug)
				msg.Collection.ContractAddress = contractAddress
			} else {
				msg.Collection.ContractAddress = msg.Collection.PrimaryAssetContracts[0].Address
			}

			_, ok := msg.CollMap[msg.Collection.Slug]
			if ok {
				msg.Collection.Ranked = true
			} else {
				msg.Collection.Ranked = false
			}

			owned = append(owned, &msg.Collection)
		case <-time.After(50 * time.Millisecond):
			return owned
		}
	}

}

func (c *collectionService) HandleCollections(request dto.CollectionRequest) []*models.Collection {
	colls := c.GetCollections(request.Owner)
	if len(colls) == 0 {
		for i := 0; i < 50; i++ {
			fmt.Println("---")
			fmt.Println("RETRY: ", i)
			fmt.Println("---")
			collections := c.GetCollections(request.Owner)
			if len(collections) > 0 {
				break
			}
		}
	}

	collsMap, _ := c.dao.NewMetaQuery().GetAllScraped()
	var tasks = &models.TaskCollections{
		Collections: colls,
		CollMap:     collsMap,
	}

	in := make(chan *models.TaskSingleCollection)
	out := make(chan *models.TaskSingleCollection)

	for i := 0; i < 15; i++ {
		go utils.Worker(in, out)
	}
	go c.Producer(in, tasks)
	collections := c.CollectionConsumer(out)
	return collections

}
