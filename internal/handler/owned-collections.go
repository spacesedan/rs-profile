package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/spacesedan/profile-tracker/internal/dto"
	"log"
	"net/http"
)

func (h *Handler) OwnedCollections(c *gin.Context) {
	var rb dto.OwnedCollectionsRequest

	if err := c.ShouldBindQuery(&rb); err != nil {
		log.Printf("Could not bind url queries: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	queries := c.Request.URL.Query()
	collections := h.CollectionService.GetOwnedCollection(queries)

	c.JSON(http.StatusOK, gin.H{
		"collections": collections,
	})

}
