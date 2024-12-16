package sqlite

import (
	"encoding/json"

	"github.com/faxa0-0/billy/plan_service/internal/models"
)

func (repo *SQLiteRepository) GetPlans() ([]*models.Plan, error) {
	query := `SELECT id, title, series, corp, subscription_fee, conn_type, speed, 
			  tasix_speed, limit_mb, additional_info, active, created_at, updated_at 
			  FROM plans`

	// Execute the query
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Slice to hold the retrieved plans
	var plans []*models.Plan

	// Iterate through the rows
	for rows.Next() {
		var plan models.Plan
		var connType string
		var speed string
		var tasixSpeed string
		var additionalInfo string

		// Scan the row into plan fields
		err := rows.Scan(
			&plan.ID, &plan.Title, &plan.Series, &plan.Corp, &plan.SubscriptionFee,
			&connType, &speed, &tasixSpeed, &plan.LimitMb, &additionalInfo,
			&plan.Active, &plan.CreatedAt, &plan.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		// Deserialize JSON fields
		if err := json.Unmarshal([]byte(connType), &plan.ConnType); err != nil {
			return nil, err
		}
		if err := json.Unmarshal([]byte(speed), &plan.Speed); err != nil {
			return nil, err
		}
		if err := json.Unmarshal([]byte(tasixSpeed), &plan.TasixSpeed); err != nil {
			return nil, err
		}
		if err := json.Unmarshal([]byte(additionalInfo), &plan.AdditionalInfo); err != nil {
			return nil, err
		}

		// Add the plan to the slice
		plans = append(plans, &plan)
	}

	// Check for errors during row iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return plans, nil
}
