package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/spacesedan/profile-tracker/internal/dto"
	"log"
	"net/http"
)

func (h *Handler) GetCollection(c *gin.Context) {
	var req dto.CollectionRequest
	var res dto.GetCollectionInformationResponse

	if err := c.BindQuery(&req); err != nil {
		log.Printf("Could not bind request queries: %v\n", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	key := "owner=" + req.Owner + "?ownedCollections"

	h.Cache.Delete(key)

	i := h.Cache.Get(key)
	if i == nil {
		log.Printf("REQUEST: %v\n", req)
		res = h.CollectionService.GetCollections(req)
		log.Printf("RES: %v\n", res)
		marshal, _ := json.Marshal(res)
		h.Cache.Set(key, string(marshal), 600)

		c.JSON(http.StatusOK, gin.H{
			"collection": res,
		})
		return
	}
	json.Unmarshal([]byte(*i), &res)

	c.JSON(http.StatusOK, gin.H{
		"collection": res,
	})

}
