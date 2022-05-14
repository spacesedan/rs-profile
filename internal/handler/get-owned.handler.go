package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spacesedan/profile-tracker/internal/dto"
	"log"
	"net/http"
)

func (h *Handler) GetOwned(c *gin.Context) {
	var req dto.AssetsWithRefresh
	var owned dto.GetOwnedResponse

	if err := c.ShouldBindQuery(&req); err != nil {
		log.Printf("Error: %v\n", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
	}

	key := "owner=" + req.Owner + "?contractAddress=" + req.ContractAddress + "?assets"

	// clear the previous value stored in redis
	if req.Refresh {
		fmt.Printf("DELETING VALUE FOR KEY: %v\n", key)
		h.Cache.Delete(key)
	}

	t := h.Cache.Get(key)

	// if "t" does not exist query mongo
	if t == nil {
		owned = h.AssetService.GetOwnedTokens(req)
		marshal, _ := json.Marshal(owned)

		h.Cache.Set(key, string(marshal), 604800)

		c.JSON(http.StatusOK, gin.H{
			"owned": owned,
		})
		return
	}
	// else unmarshal the value stored in redis and return it to the user
	json.Unmarshal([]byte(*t), &owned)
	c.JSON(http.StatusOK, gin.H{
		"owned": owned,
	})
}
