package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

func (h *Handler) version(c *gin.Context) {
	ctx := c.Request.Context()
	span, ctx := opentracing.StartSpanFromContext(ctx, "serve.Version")
	defer span.Finish()

	logger := h.WrappedLogger(ctx)

	logger.Infow("serving version", "version", apiVersionNumber)
	c.JSON(http.StatusOK, gin.H{"version": apiVersionNumber})
}
