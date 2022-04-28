package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/spacesedan/profile-tracker/internal/dto"
	"net/http"
)

func (h *Handler) Collections(c *gin.Context) {
	var rb dto.CollectionRequest

	err := c.BindJSON(&rb)
	if err != nil {
		return
	}

	collections := h.CollectionService.HandleCollections(rb)

	c.JSON(http.StatusOK, gin.H{
		"collections": collections,
	})
}
