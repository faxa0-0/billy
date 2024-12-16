package sqlite

import (
	"fmt"

	"github.com/faxa0-0/billy/plan_service/internal/models"
)

func (repo *SQLiteRepository) CreatePlan(plan models.Plan) (*models.CreateResult, error) {
	fmt.Println("repo.CreatePlan")
	return nil, nil
}
