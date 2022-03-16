package handler

import (
	"context"
	"strconv"
	"time"

	"github.com/blendle/zapdriver"
	pb "github.com/dmartzol/goapi/internal/proto"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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
		LogHandler(h.SugaredLogger),
		gin.Recovery(),
		//h.AuthMiddleware,
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

func LogHandler(logger *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		//fields := zapdriver.Operation(
		//strconv.Itoa(int(time.Now().UnixNano())),
		//"gateway",
		//true,
		//false,
		//)
		//logger = logger.Desugar().With(fields).Sugar()

		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		//if _, ok := skipPaths[path]; !ok {
		end := time.Now()
		end = end.UTC() // always use UTC
		latency := end.Sub(start)

		if len(c.Errors) > 0 {
			// Append error field if this is an erroneous request.
			for _, e := range c.Errors.Errors() {
				logger.Error(e)
			}
		} else {
			fields := []zapcore.Field{
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", c.ClientIP()),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.Duration("latency", latency),
			}
			//fields = append(fields, zap.String("time", end.Format(time.RFC3339)))
			logger.Desugar().With(fields...).Info(strconv.Itoa(c.Writer.Status()))
			//logger = logger.With(fields)
			//logger.Infow(path, "key", "value")
			//}
		}

	}
}
