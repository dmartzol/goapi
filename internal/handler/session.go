package handler

import (
	"net/http"

	"github.com/dmartzol/goapi/goapi"
	"github.com/dmartzol/goapi/pkg/httputils"
)

const (
	// sessionLength represents the duration(in seconds) a session will be valid for
	sessionLength = 345600
)

func (h *Handler) GetSession(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie(CookieName)
	if err != nil {
		h.Errorw("could not fetch cookie", "cookie", CookieName, "error", err)
		httputils.RespondJSONError(w, "", http.StatusInternalServerError)
		return
	}
	// s, err := h.db.SessionFromToken(c.Value)
	if err != nil {
		h.Errorw("could not fetch session from token", "token", c.Value, "error", err)
		httputils.RespondJSONError(w, "", http.StatusUnauthorized)
		return
	}
	// httputils.RespondJSON(w, s.View(nil))
}

func (h *Handler) CreateSession(w http.ResponseWriter, r *http.Request) {
	var credentials goapi.LoginCredentials
	err := h.Unmarshal(r, &credentials)
	if err != nil {
		h.Errorw("could not unmarshal", "error", err)
		httputils.RespondJSONError(w, "", http.StatusInternalServerError)
		return
	}

	// fetching account with credentials(errors reurned should be purposedly broad)
	// a, err := h.db.AccountWithCredentials(credentials.Email, credentials.Password)
	if err != nil {
		h.Errorw("could not fetch account", "email", credentials.Email)
		httputils.RespondJSONError(w, "", http.StatusUnauthorized)
		return
	}
	credentials.Password = ""

	// create session and cookie
	// s, err := h.db.CreateSession(a.ID)
	// if err != nil {
	// 	h.Errorw("could not create session", "account", a.ID)
	// 	httputils.RespondJSONError(w, "", http.StatusInternalServerError)
	// 	return
	// }
	cookie := &http.Cookie{
		Name: CookieName,
		// Value:  s.Token,
		MaxAge: sessionLength,
	}
	http.SetCookie(w, cookie)
	// httputils.RespondJSON(w, s.View(nil))
}

func (h *Handler) ExpireSession(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie(CookieName)
	if err != nil {
		h.Errorw("could not fetch cookie", "cookie", CookieName, "error", err)
		httputils.RespondJSONError(w, "", http.StatusInternalServerError)
		return
	}
	// session, err := h.db.ExpireSessionFromToken(c.Value)
	if err != nil {
		h.Errorw("could not expire session", "token", c.Value, "error", err)
		httputils.RespondJSONError(w, "", http.StatusInternalServerError)
		return
	}
	c = &http.Cookie{
		Name:   CookieName,
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)
	// httputils.RespondJSON(w, session.View(nil))
}
