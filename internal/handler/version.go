package handler

import (
	"fmt"
	"net/http"

	"github.com/dmartzol/api-template/pkg/httpresponse"
)

func (h *Handler) Version(w http.ResponseWriter, r *http.Request) {
	h.Infow("serving service version", "version", apiVersionNumber)
	httpresponse.RespondJSON(w, fmt.Sprintf("version %s", apiVersionNumber))
}
