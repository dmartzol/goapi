package handler

import (
	"log"
	"net/http"

	"github.com/dmartzol/api-template/internal/storage/postgres"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

const (
	apiVersionNumber = "0.0.1"
	CookieName       = "API-Template-Cookie"
	idQueryParameter = "id"
)

type Handler struct {
	Router *mux.Router
	db     *postgres.DB
}

func NewHandler(router *mux.Router, db *postgres.DB) (*Handler, error) {
	return &Handler{
		Router: router,
		db:     db,
	}, nil
}

func (h *Handler) InitializeRoutes() {
	h.Router = h.Router.PathPrefix("/v1").Subrouter()

	h.Router.Use(
		middleware.Logger,
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
	log.Printf("listening and serving on %s", addr)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		// Debug: true,
	})
	log.Fatal(http.ListenAndServe(addr, c.Handler(h.Router)))
}
