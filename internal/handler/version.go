package handler

import (
	"fmt"
	"net/http"
)

func (h *handler) Version(w http.ResponseWriter, r *http.Request) {
	h.Infow("serving service version", "version", apiVersionNumber)
	fmt.Fprintf(w, "version %s", apiVersionNumber)
}
