package postgres

import (
	"database/sql"
	"time"

	"github.com/faxa0-0/billy/user_service/internal/models"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (repo *PostgresUserRepository) Create(user *models.User) error { return nil }
func (repo *PostgresUserRepository) FindByID(id string) (*models.User, error) {
	return &models.User{
		ID:              "123",
		Name:            "John Doe",
		Login:           "jdoe",
		PaymentAcc:      "acc-456",
		ConnType:        "fttx",
		Balance:         11000000, // e.g., balance in cents
		WriteOffDate:    time.Now(),
		Active:          true,
		PlanTitle:       "Premium Plan",
		PlanSeries:      "A",
		PlanSubsFee:     9900000, // e.g., subscription fee in cents
		LastPaymentSum:  5600000,
		LastPaymentDate: time.Now(),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}, nil
}
func (repo *PostgresUserRepository) FindAll() ([]models.User, error)           { return nil, nil }
func (repo *PostgresUserRepository) Update(id string, user *models.User) error { return nil }
func (repo *PostgresUserRepository) Delete(id string) error                    { return nil }
