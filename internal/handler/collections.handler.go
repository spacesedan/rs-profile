package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/spacesedan/profile-tracker/internal/dto"
	"log"
	"net/http"
)

func (h *Handler) Collections(c *gin.Context) {
	var rb dto.CollectionRequest
	var collections dto.GetCollectionInformationResponse

	if err := c.BindQuery(&rb); err != nil {
		log.Printf("Could not parse queries: %v\n", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
	}

	key := "owner=" + rb.Owner + "?ownedCollectionSlugs"

	if rb.Refresh {
		h.Cache.Delete(key)
	}

	v := h.Cache.Get(key)
	if v == nil {
		collections = h.CollectionService.GetCollections(rb)

		marshal, _ := json.Marshal(collections)

		h.Cache.Set(key, string(marshal), 604800*3)
		c.JSON(http.StatusOK, gin.H{
			"collections": collections,
		})
		return
	}
	json.Unmarshal([]byte(*v), &collections)

	c.JSON(http.StatusOK, gin.H{
		"collections": collections,
	})

}
