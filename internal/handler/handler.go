package handler

import (
	"log"
	"net/http"

	"github.com/dmartzol/api-template/internal/storage/postgres"
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
	Router *mux.Router
	db     *postgres.DB
}

func NewHandler(router *mux.Router, db *postgres.DB, logger *zap.SugaredLogger) *Handler {
	return &Handler{
		SugaredLogger: logger,
		Router:        router,
		db:            db,
	}
}

func (h *Handler) InitializeRoutes() {
	h.Router = h.Router.PathPrefix("/v1").Subrouter()

	h.Router.Use(
		middleware.Recoverer,
		h.AuthMiddleware,
	)

	h.Router.HandleFunc("/version", h.Version).Methods("GET")

	// sessions
	// see: https://stackoverflow.com/questions/7140074/restfully-design-login-or-register-resources
	h.Router.HandleFunc("/sessions", h.CreateSession).Methods("POST")
	h.Router.HandleFunc("/sessions", h.GetSession).Methods("GET")
	h.Router.HandleFunc("/sessions", h.ExpireSession).Methods("DELETE")
}

func (h *Handler) Run(addr string) {
	h.Infof("listening and serving on %s", addr)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		// Debug: true,
	})
	log.Fatal(http.ListenAndServe(addr, c.Handler(h.Router)))
}