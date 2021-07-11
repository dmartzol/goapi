package handler

import (
	"log"
	"net/http"

	models "github.com/dmartzol/api-template/internal/model"
	"github.com/dmartzol/api-template/pkg/httpresponse"
)

const (
	// sessionLength represents the duration(in seconds) a session will be valid for
	sessionLength = 345600
)

func (h *handler) GetSession(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie(CookieName)
	if err != nil {
		log.Printf("%+v", err)
		httpresponse.RespondJSONError(w, "", http.StatusInternalServerError)
		return
	}
	s, err := h.db.SessionFromToken(c.Value)
	if err != nil {
		log.Printf("%+v", err)
		httpresponse.RespondJSONError(w, "", http.StatusUnauthorized)
		return
	}
	httpresponse.RespondJSON(w, s.View(nil))
}

func (h *handler) CreateSession(w http.ResponseWriter, r *http.Request) {
	var credentials models.LoginCredentials
	err := httpresponse.Unmarshal(r, &credentials)
	if err != nil {
		log.Printf("Unmarshal: %+v", err)
		httpresponse.RespondJSONError(w, "", http.StatusInternalServerError)
		return
	}

	// fetching account with credentials(errors reurned should be purposedly broad)
	a, err := h.db.AccountWithCredentials(credentials.Email, credentials.Password)
	if err != nil {
		log.Printf("%+v", err)
		httpresponse.RespondJSONError(w, "", http.StatusUnauthorized)
		return
	}
	credentials.Password = ""

	// create session and cookie
	s, err := h.db.CreateSession(a.ID)
	if err != nil {
		log.Printf("%+v", err)
		httpresponse.RespondJSONError(w, "", http.StatusInternalServerError)
		return
	}
	cookie := &http.Cookie{
		Name:   CookieName,
		Value:  s.Token,
		MaxAge: sessionLength,
	}
	http.SetCookie(w, cookie)
	httpresponse.RespondJSON(w, s.View(nil))
}

func (h *handler) ExpireSession(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie(CookieName)
	if err != nil {
		log.Printf("%+v", err)
		httpresponse.RespondJSONError(w, "", http.StatusInternalServerError)
		return
	}
	session, err := h.db.ExpireSessionFromToken(c.Value)
	if err != nil {
		log.Printf("ExpireSession - ERROR expiring session: %+v", err)
		httpresponse.RespondJSONError(w, "", http.StatusInternalServerError)
		return
	}
	c = &http.Cookie{
		Name:   CookieName,
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)
	httpresponse.RespondJSON(w, session.View(nil))
}
