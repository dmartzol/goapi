package handler

import (
	"fmt"
	"net/http"

	"github.com/dmartzol/goapi/pkg/httputils"
)

func (h *Handler) Version(w http.ResponseWriter, r *http.Request) {
	h.Infow("serving service version", "version", apiVersionNumber)
	httputils.RespondJSON(w, fmt.Sprintf("version %s", apiVersionNumber))
}
