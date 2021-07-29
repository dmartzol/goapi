package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/dmartzol/api-template/internal/api"
	"github.com/dmartzol/api-template/internal/model"
	pb "github.com/dmartzol/api-template/internal/protos"
	"github.com/dmartzol/api-template/pkg/httputils"
)

func (h *Handler) createAccount(w http.ResponseWriter, r *http.Request) {
	var req api.CreateAccountRequest
	err := httputils.Unmarshal(r, &req)
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
	addAccountReq := pb.AddAccountMessage{
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	aM, err := h.Accounts.AddAccount(ctx, &addAccountReq)
	if err != nil {
		h.Errorw("failed to create account", "error", err)
		httputils.RespondJSONError(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	h.Debugf("aM: %v", aM)
	a, err := model.MarshallAccount(aM)
	h.Debugf("a: %v", a)
	httputils.RespondJSONError(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}
