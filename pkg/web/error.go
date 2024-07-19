package web

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func Error(w http.ResponseWriter, err string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := ErrorResponse{
		StatusCode: statusCode,
		Message:    err,
	}

	json.NewEncoder(w).Encode(response)

}
