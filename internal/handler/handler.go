package handler

import (
	pb "github.com/dmartzol/goapi/internal/proto"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	apiVersionNumber = "0.0.1"
	CookieName       = "goapi-Cookie"
	idQueryParameter = "id"
)

type Handler struct {
	*zap.SugaredLogger
	Accounts      pb.AccountsClient
	Router        *gin.Engine
	LogRawRequest bool
}

func New(ac pb.AccountsClient, logger *zap.SugaredLogger, logRawRequest bool) (*Handler, error) {
	h := Handler{
		Router:        gin.Default(),
		Accounts:      ac,
		SugaredLogger: logger,
		LogRawRequest: logRawRequest,
	}

	return &h, nil
}

func (h *Handler) InitializeRoutes() {

	//h.Router.Use(
	//middleware.Logger,
	//middleware.Recoverer,
	//h.AuthMiddleware,
	//)

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
