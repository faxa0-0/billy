package handler

import (
	"encoding/json"
	"net/http"

	"github.com/faxa0-0/billy/plan_service/internal/models"
	"github.com/faxa0-0/billy/plan_service/internal/utils"
)

func (h *PlanHandler) CreatePlan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newPlan models.Plan

	if err := json.NewDecoder(r.Body).Decode(&newPlan); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	createdPlan, err := h.planService.CreatePlan(newPlan)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Error creating plan")
		return
	}

	response := models.Response{
		StatusCode: http.StatusCreated,
		Status:     "success",
		Message:    "Plan created successfully",
		Data:       createdPlan,
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
