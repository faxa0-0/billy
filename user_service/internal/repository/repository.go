package repository

import "github.com/faxa0-0/billy/plan_service/internal/models"

type PlanRepository interface {
	Close() error
	CreatePlan(plan models.Plan) (*models.CreatedResponse, error)
	GetPlans() ([]*models.Plan, error)
}
