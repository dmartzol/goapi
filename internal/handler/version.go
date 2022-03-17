package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

func (h *Handler) version(c *gin.Context) {
	ctx := c.Request.Context()
	logger := h.WrappedLogger(ctx)

	var span opentracing.Span
	span, ctx = opentracing.StartSpanFromContext(ctx, "serve.Version")
	defer span.Finish()

	another(ctx)

	logger.Infow("serving version", "version", apiVersionNumber)
	c.JSON(http.StatusOK, gin.H{"version": apiVersionNumber})
}

func another(ctx context.Context) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "serve.Version")
	defer span.Finish()
}
