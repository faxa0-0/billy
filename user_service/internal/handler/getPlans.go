package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/faxa0-0/billy/plan_service/internal/models"
	"github.com/faxa0-0/billy/plan_service/internal/utils"
)

func (h *PlanHandler) GetPlans(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	plans, err := h.planService.GetPlans()
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Error finding plan")
		return
	}

	response := models.Response{
		StatusCode: http.StatusOK,
		Status:     "success",
		Message:    fmt.Sprintf("%d plans found", len(plans)),
		Data:       plans,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
