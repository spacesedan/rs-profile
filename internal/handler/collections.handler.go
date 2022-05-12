package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/spacesedan/profile-tracker/internal/dto"
	"log"
	"net/http"
)

func (h *Handler) Collections(c *gin.Context) {
	var rb dto.CollectionRequest

	if err := c.ShouldBindQuery(&rb); err != nil {
		log.Printf("Could not parse queries: %v\n", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
	}

	collections := h.CollectionService.GetCollections(rb)

	c.JSON(http.StatusOK, gin.H{
		"collections": collections,
	})
}
