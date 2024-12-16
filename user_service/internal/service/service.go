package service

import (
	"fmt"

	"github.com/faxa0-0/billy/plan_service/internal/models"
	"github.com/faxa0-0/billy/plan_service/internal/repository"
)

type PlanService struct {
	repo repository.PlanRepository
}

func NewPlanService(repo repository.PlanRepository) *PlanService {
	return &PlanService{repo: repo}
}

func (s *PlanService) CreatePlan(newPlan models.Plan) (*models.CreatedResponse, error) {
	if newPlan.Title == "" {
		return nil, fmt.Errorf("plan title is required")
	}

	createdPlan, err := s.repo.CreatePlan(newPlan)
	if err != nil {
		return nil, fmt.Errorf("error saving plan: %v", err)
	}

	return createdPlan, nil
}

func (s *PlanService) GetPlans() ([]*models.Plan, error) {
	plans, err := s.repo.GetPlans()
	if err != nil {
		return nil, fmt.Errorf("error finding plans: %v", err)
	}

	return plans, nil
}
