package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/spacesedan/profile-tracker/internal/dto"
	"log"
	"net/http"
)

func (h *Handler) GetOwned(c *gin.Context) {
	var req dto.AssetsWithCursorRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		log.Printf("Error: %v\n", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	tokenIds := h.AssetService.GetOwnedTokenIds(req)

	c.JSON(http.StatusOK, gin.H{
		"owned": tokenIds,
	})
}
