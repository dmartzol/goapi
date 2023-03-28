package api

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/blendle/zapdriver"
	pb "github.com/dmartzol/goapi/internal/proto"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

const (
	apiVersionNumber = "0.0.1"
	CookieName       = "goapi-Cookie"
	idQueryParameter = "id"
)

type Handler struct {
	projectID string
	*zap.SugaredLogger
	Accounts      pb.AccountsClient
	Router        *gin.Engine
	LogRawRequest bool
}

func New(ac pb.AccountsClient, logger *zap.SugaredLogger, logRawRequest bool) (*Handler, error) {
	h := Handler{
		Router:        gin.New(),
		Accounts:      ac,
		SugaredLogger: logger,
		LogRawRequest: logRawRequest,
	}
	return &h, nil
}

func (h *Handler) InitializeRoutes() {
	h.Router.Use(
		gin.Recovery(),
		h.AuthMiddleware,
	)

	//log.Infow("Test Log", zap.String("trace", "projects/gcp-project-id/traces/"+span.SpanContext().TraceID.String()))
	v1 := h.Router.Group("/v1")
	{
		v1.GET("/version", h.ginVersion)
		v1.POST("/accounts", h.createAccount)
		// sessions
		v1.POST("/sessions", h.createSession)
		v1.GET("/sessions", h.getSession)
		// see: https://stackoverflow.com/questions/7140074/restfully-design-login-or-register-resources
		v1.DELETE("/sessions", h.expireSession)
	}
}

func (h *Handler) WrappedLogger(ctx context.Context) *zap.SugaredLogger {
	sc := trace.SpanContextFromContext(ctx)
	fields := zapdriver.TraceContext(sc.TraceID().String(), sc.SpanID().String(), sc.IsSampled(), h.projectID)
	logger := h.SugaredLogger.Desugar().With(fields...)
	return logger.Sugar()
}

func (h *Handler) Unmarshal(c *gin.Context, iface interface{}) error {
	b, err := c.GetRawData()
	if err != nil {
		return fmt.Errorf("unable to get raw data: %w", err)
	}
	if h.LogRawRequest {
		h.Infof("payload: %s", b)
	}
	err = json.Unmarshal(b, &iface)
	if err != nil {
		h.Errorf("json.Unmarshal %+v", err)
		return err
	}
	return nil
}
