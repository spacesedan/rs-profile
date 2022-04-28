package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/spacesedan/profile-tracker/internal/dto"
	"log"
	"net/http"
)

func (h *Handler) Assets(c *gin.Context) {
	var rb dto.AssetRequest

	if err := c.ShouldBindQuery(&rb); err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	log.Println(rb)
	assets := h.AssetService.HandleAssets(rb)
	c.JSON(http.StatusOK, gin.H{
		"assets": assets,
	})
}
