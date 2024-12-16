package handler

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/faxa0-0/billy/plan_service/internal/service"
)

type PlanHandler struct {
	taskService *service.PlanService
}

func NewPlanHandler(service *service.PlanService) *PlanHandler {
	return &PlanHandler{taskService: service}
}
func (handler *PlanHandler) PlansHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("GET plans"))
	case http.MethodPost:
		w.Write([]byte("POST plans"))
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
func (handler *PlanHandler) SinglePlanHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/plans/"):]

	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}
	var re = regexp.MustCompile(`^[0-9]+$`)

	if !re.MatchString(id) {
		http.Error(w, "bad id", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "GET plan with ID: %s", id)
	case http.MethodPatch:
		fmt.Fprintf(w, "PATCH plan with ID: %s", id)
	case http.MethodDelete:
		fmt.Fprintf(w, "DELETE plan with ID: %s", id)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
