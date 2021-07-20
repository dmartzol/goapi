package handler

import (
	"net/http"

	"github.com/dmartzol/api-template/internal/api"
	"github.com/dmartzol/api-template/internal/model"
	"github.com/dmartzol/api-template/pkg/httputils"
)

func (h *Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var req api.CreateAccountRequest
	err := httputils.Unmarshal(r, &req)
	if err != nil {
		h.Errorw("could not unmarshal", "error", err)
		httputils.RespondJSONError(w, "", http.StatusInternalServerError)
		return
	}
	a := model.Account{}
	httputils.RespondJSON(w, a.View(nil))
}
