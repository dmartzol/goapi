package handler

import "github.com/dmartzol/api-template/internal/storage/postgres"

const (
	apiVersionNumber = "0.0.1"
	CookieName       = "API-Template-Cookie"
	idQueryParameter = "id"
)

// Handler represents something
type Handler struct {
	db *postgres.DB
}

func NewHandler(db *postgres.DB) (*Handler, error) {
	return &Handler{db}, nil
}
