package handler

import (
	"context"
	"database/sql"
	"errors"
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
			h.Errorw("could not fetch cookie", "cookie", CookieName, "error", err)
			httpresponse.RespondJSONError(w, "", http.StatusUnauthorized)
			return
		}
		s, err := h.db.UpdateSession(c.Value)
		if err != nil {
			h.Errorw("could not update session", "token", c.Value, "error", err)
			if err == sql.ErrNoRows {
				h.Errorw("unable to find session", "error", err)
				httpresponse.RespondJSONError(w, "", http.StatusUnauthorized)
				return
			}
			if errors.Is(err, postgres.ErrExpiredSession) {
				h.Errorw("expired session", "error", err)
				httpresponse.RespondJSONError(w, "", http.StatusUnauthorized)
				return
			}
			httpresponse.RespondJSONError(w, "", http.StatusInternalServerError)
			return
		}

		// Setting up context
		ctx := r.Context()
		ctx = context.WithValue(ctx, contextRequesterAccountIDKey, s.AccountID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
