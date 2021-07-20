package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/dmartzol/api-template/internal/api"
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
	a := pb.CreateAccountMessage{
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	b, err := h.CreateAccount(ctx, &a)
	if err != nil {
		h.Errorw("failed to create account", "error", err)
		httputils.RespondJSONError(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	h.Debugf("b: %v", b)
	httputils.RespondJSONError(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}
