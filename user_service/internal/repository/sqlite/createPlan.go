package sqlite

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/faxa0-0/billy/plan_service/internal/models"
)

func (repo *SQLiteRepository) CreatePlan(plan models.Plan) (*models.CreatedResponse, error) {
	connTypeJSON, err := json.Marshal(plan.ConnType)
	if err != nil {
		return nil, fmt.Errorf("error encoding conn_type: %w", err)
	}

	speedJSON, err := json.Marshal(plan.Speed)
	if err != nil {
		return nil, fmt.Errorf("error encoding speed: %w", err)
	}

	tasixSpeedJSON, err := json.Marshal(plan.TasixSpeed)
	if err != nil {
		return nil, fmt.Errorf("error encoding tasix_speed: %w", err)
	}

	additionalInfo := fmt.Sprintf(`{"ru": "%s", "en": "%s", "uzb": "%s"}`, plan.AdditionalInfo.Ru, plan.AdditionalInfo.En, plan.AdditionalInfo.Uzb)

	createdAt := time.Now()
	updatedAt := createdAt

	query := `INSERT INTO plans (title, series, corp, subscription_fee, conn_type, speed, tasix_speed, limit_mb, additional_info, active, created_at, updated_at)
			  VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := repo.db.Exec(query, plan.Title, plan.Series, plan.Corp, plan.SubscriptionFee, string(connTypeJSON), string(speedJSON), string(tasixSpeedJSON), plan.LimitMb, additionalInfo, plan.Active, createdAt, updatedAt)
	if err != nil {
		return nil, fmt.Errorf("error saving plan: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error retrieving last insert ID: %w", err)
	}

	return &models.CreatedResponse{ID: int(id), Title: plan.Title}, nil
}
