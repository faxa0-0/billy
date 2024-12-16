package utils

import (
	"encoding/json"
	"net/http"

	"github.com/faxa0-0/billy/plan_service/internal/models"
)

func WriteErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	response := models.Response{
		StatusCode: statusCode,
		Status:     "failed",
		Message:    message,
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
