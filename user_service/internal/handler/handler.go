package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/faxa0-0/billy/user_service/internal/models"
	"github.com/faxa0-0/billy/user_service/internal/service"
	"github.com/faxa0-0/billy/user_service/pkg/response"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) HandleUserRequests(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		id := strings.TrimPrefix(r.URL.Path, "/users/")
		if id == "" || id == "/users" {
			h.GetUsers(w, r)
		} else {
			h.GetUserByID(w, r, id)
		}
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
func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request, id string) {
	user, err := h.service.GetUser(id)
	if err != nil {
		response.ErrorResponse(w, "User not found", http.StatusNotFound)
		return
	}
	response.SuccessResponse(w, user, "user found successfully")
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.ListUsers()
	if err != nil || len(users) == 0 {
		response.ErrorResponse(w, "Users not found", http.StatusNotFound)
		return
	}

	response.SuccessResponse(w, users, "users found successfully")
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.ErrorResponse(w, "Bad request", http.StatusBadRequest)
		return
	}
	id, err := h.service.CreateUser(&user)
	if err != nil {
		response.ErrorResponse(w, "Server error", http.StatusInternalServerError)
		return
	}
	response.SuccessResponse(w, struct {
		ID int `json:"id"`
	}{ID: id}, "User created successfully")
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/users/")

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.ErrorResponse(w, "Bad request", http.StatusBadRequest)
		return
	}

	err := h.service.UpdateUser(id, &user)
	if err != nil {
		response.ErrorResponse(w, "Failed to update user info", http.StatusNotFound)
		return
	}
	response.SuccessResponse(w, nil, "user info updated successfully")
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/users/")

	err := h.service.DeleteUser(id)
	if err != nil {
		response.ErrorResponse(w, "Failed to delete user", http.StatusNotFound)
		return
	}

	response.SuccessResponse(w, nil, "user deleted successfully")
}
