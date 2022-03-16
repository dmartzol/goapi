package handler

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ErrExpiredSession error
)

func (h *Handler) AuthMiddleware(c *gin.Context) {

	publicRoutes := map[string]string{
		"/v1/version":  "GET",
		"/v1/sessions": "POST",
		"/v1/accounts": "POST",
	}
	method, ok := publicRoutes[c.Request.URL.Path]
	if ok && method == c.Request.Method {
		c.Next()
		return
	}
	cookie, err := c.Cookie(CookieName)
	if err != nil {
		h.Errorw("could not fetch cookie", "cookie", CookieName, "error", err)
		return
	}
	//s, err := h.db.UpdateSession(cookie)
	if err != nil {
		h.Errorw("could not update session", "token", cookie, "error", err)
		if err == sql.ErrNoRows {
			h.Errorw("unable to find session", "error", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, ErrExpiredSession) {
			h.Errorw("expired session", "error", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Setting up context
	//c.Set(contextRequesterAccountIDKey, s.AccountID)
}
