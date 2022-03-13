package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ginVersion(c *gin.Context) {
	h.Infow("serving service version", "version", apiVersionNumber)
	c.JSON(http.StatusOK, gin.H{"version": apiVersionNumber})
}
