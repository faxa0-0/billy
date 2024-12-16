package service

import "github.com/faxa0-0/billy/plan_service/internal/repository"

type PlanService struct {
	repo repository.PlanRepository
}

func NewPlanService(repo repository.PlanRepository) *PlanService {
	return &PlanService{repo: repo}
}
