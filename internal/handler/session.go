package handler

import (
	"net/http"

	"github.com/dmartzol/goapi/goapi"
	"github.com/gin-gonic/gin"
)

const (
	// sessionLength represents the duration(in seconds) a session will be valid for
	sessionLength = 345600
)

func (h *Handler) getSession(c *gin.Context) {
	_, err := c.Cookie(CookieName)
	if err != nil {
		h.Errorw("could not fetch cookie", "cookie", CookieName, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// s, err := h.db.SessionFromToken(c.Value)
	if err != nil {
		h.Errorw("could not fetch session from token", "token", c.Value, "error", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	// httputils.RespondJSON(w, s.View(nil))
}

func (h *Handler) createSession(c *gin.Context) {
	var credentials goapi.LoginCredentials
	err := h.Unmarshal(c, &credentials)
	if err != nil {
		h.Errorw("could not unmarshal", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// fetching account with credentials(errors reurned should be purposedly broad)
	// a, err := h.db.AccountWithCredentials(credentials.Email, credentials.Password)
	if err != nil {
		h.Errorw("could not fetch account", "email", credentials.Email)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	credentials.Password = ""

	// create session and cookie
	//s, err := h.db.CreateSession(a.ID)
	//if err != nil {
	//h.Errorw("could not create session", "account", a.ID)
	//httputils.RespondJSONError(w, "", http.StatusInternalServerError)
	//return
	//}
	cookie := &http.Cookie{
		Name: CookieName,
		// Value:  s.Token,
		MaxAge: sessionLength,
	}
	c.SetCookie(
		CookieName,
		"",
		//s.Token,
		cookie.MaxAge,
		cookie.Path,
		cookie.Domain,
		cookie.Secure,
		cookie.HttpOnly,
	)
	//c.JSON(http.StatusOK, s.View(nil))
}

func (h *Handler) expireSession(c *gin.Context) {
	_, err := c.Cookie(CookieName)
	if err != nil {
		h.Errorw("could not fetch cookie", "cookie", CookieName, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//session, err := h.Sessions.ExpireSessionFromToken(c.Value)
	if err != nil {
		h.Errorw("could not expire session", "token", c.Value, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie(CookieName, "", -1, "", "", true, true)
	// httputils.RespondJSON(w, session.View(nil))
	//c.JSON(http.StatusOK, session.View(nil))
}
