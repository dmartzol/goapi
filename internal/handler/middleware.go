package handler

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"

	"github.com/dmartzol/api-template/internal/storage/postgres"
	"github.com/dmartzol/api-template/pkg/httpresponse"
)

func (h *handler) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		publicRoutes := map[string]string{
			"/v1/version":  "GET",
			"/v1/sessions": "POST",
			"/v1/accounts": "POST",
		}
		method, in := publicRoutes[r.RequestURI]
		if in && method == r.Method {
			next.ServeHTTP(w, r)
			return
		}
		c, err := r.Cookie(CookieName)
		if err != nil {
			log.Printf("AuthMiddleware ERROR getting cookie: %+v", err)
			httpresponse.RespondJSONError(w, "", http.StatusUnauthorized)
			return
		}
		s, err := h.db.UpdateSession(c.Value)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Printf("AuthMiddleware ERROR unable to find session %s: %+v", c.Value, err)
				httpresponse.RespondJSONError(w, "", http.StatusUnauthorized)
				return
			}
			if errors.Is(err, postgres.ErrExpiredSession) {
				h.Errorw("expired session", "session", c.Value, "error", err)
				httpresponse.RespondJSONError(w, "", http.StatusUnauthorized)
				return
			}
			log.Printf("AuthMiddleware ERROR for session %s: %+v", c.Value, err)
			httpresponse.RespondJSONError(w, "", http.StatusInternalServerError)
			return
		}

		// Setting up context
		ctx := r.Context()
		ctx = context.WithValue(ctx, contextRequesterAccountIDKey, s.AccountID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
