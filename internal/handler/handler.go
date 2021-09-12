package handler

import (
	"net/http"

	pb "github.com/dmartzol/api-template/internal/protos"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.uber.org/zap"
)

const (
	apiVersionNumber = "0.0.1"
	CookieName       = "API-Template-Cookie"
	idQueryParameter = "id"
)

type Handler struct {
	*zap.SugaredLogger
	Accounts pb.AccountsClient
	Router   *mux.Router
}

func New(ac pb.AccountsClient, logger *zap.SugaredLogger) (*Handler, error) {
	h := Handler{
		Accounts:      ac,
		Router:        mux.NewRouter(),
		SugaredLogger: logger,
	}

	return &h, nil
}

func (h *Handler) InitializeRoutes() {
	h.Router = h.Router.PathPrefix("/v1").Subrouter()

	h.Router.Use(
		middleware.Recoverer,
		h.AuthMiddleware,
	)

	h.Router.HandleFunc("/version", h.Version).Methods("GET")

	// accounts
	h.Router.HandleFunc("/accounts", h.createAccount).Methods("POST")

	// sessions
	// see: https://stackoverflow.com/questions/7140074/restfully-design-login-or-register-resources
	h.Router.HandleFunc("/sessions", h.CreateSession).Methods("POST")
	h.Router.HandleFunc("/sessions", h.GetSession).Methods("GET")
	h.Router.HandleFunc("/sessions", h.ExpireSession).Methods("DELETE")
}

func (h *Handler) Run(addr string) error {
	h.Infof("listening and serving on %s", addr)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		// Debug: true,
	})
	return http.ListenAndServe(addr, c.Handler(h.Router))
}
