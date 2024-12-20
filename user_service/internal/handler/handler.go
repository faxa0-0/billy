package handler

import (
	"net/http"

	"github.com/faxa0-0/billy/user_service/internal/service"
	"github.com/faxa0-0/billy/user_service/pkg/response"
)

type UserHandler struct {
	service *service.UserService
}

func (h *UserHandler) HandleUserRequests(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.GetUser(w, r)
	case "POST":
		h.CreateUser(w, r)
	case "PUT":
		h.UpdateUser(w, r)
	case "DELETE":
		h.DeleteUser(w, r)
	default:
		response.ErrorResponse(w, "method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	user, err := h.service.GetUser(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	response.SuccessResponse(w, user, "user found successfully")
}
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {}
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {}
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {}
