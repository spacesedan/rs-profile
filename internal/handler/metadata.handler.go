package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/spacesedan/profile-tracker/internal/dto"
	"net/http"
)

func (h *Handler) Metadata(c *gin.Context) {
	var rb dto.MetadataRequest

	err := c.ShouldBindUri(&rb)
	if err != nil {
		return
	}

	msg := h.MetadataService.HandleMetadata(rb)

	c.JSON(http.StatusOK, gin.H{
		"message": msg,
	})

}
