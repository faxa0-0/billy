package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessResponse(w http.ResponseWriter, data interface{}, message string) {
	res := Response{
		Message: message,
		Data:    data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func ErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	res := Response{
		Message: message,
		Data:    nil,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(res)
}
