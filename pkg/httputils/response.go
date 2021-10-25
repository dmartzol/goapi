package httputils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JSONError struct {
	Error      string
	StatusCode int
}

func RespondText(w http.ResponseWriter, text string, code int) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	fmt.Fprint(w, text)
}

func RespondJSON(w http.ResponseWriter, object interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	json.NewEncoder(w).Encode(object)
}

func RespondJSONError(w http.ResponseWriter, errorMessage string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	if errorMessage == "" {
		errorMessage = http.StatusText(code)
	}
	e := JSONError{
		Error:      errorMessage,
		StatusCode: code,
	}
	json.NewEncoder(w).Encode(e)
}
