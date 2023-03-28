package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ginVersion(c *gin.Context) {
	ctx := c.Request.Context()
	logger := h.WrappedLogger(ctx)
	logger.Infow("serving version", "version", apiVersionNumber)
	c.JSON(http.StatusOK, gin.H{"version": apiVersionNumber})
}
