package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/dmartzol/goapi/internal/api"
	"github.com/dmartzol/goapi/internal/proto"
	pb "github.com/dmartzol/goapi/internal/proto"
	"github.com/dmartzol/goapi/pkg/httputils"
)

func (h *Handler) createAccount(w http.ResponseWriter, r *http.Request) {
	var req api.CreateAccountRequest
	err := h.Unmarshal(r, &req)
	if err != nil {
		h.Errorw("could not unmarshal", "error", err)
		httputils.RespondJSONError(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if err := req.Validate(); err != nil {
		h.Errorw("failed to validate create account request", "error", err)
		httputils.RespondJSONError(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	addAccountMessage := pb.AddAccountMessage{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	pbAccount, err := h.Accounts.AddAccount(ctx, &addAccountMessage)
	if err != nil {
		h.Errorw("failed to create account", "error", err)
		httputils.RespondJSONError(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	h.Debugf("pbAccount: %v", pbAccount)
	a, err := proto.CoreAccount(pbAccount)
	if err != nil {
		h.Errorw("failed to marshall account", "error", err)
		httputils.RespondJSONError(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	h.Debugf("a: %v", a)
	httputils.RespondJSONError(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}
