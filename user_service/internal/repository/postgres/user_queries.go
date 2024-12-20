package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/faxa0-0/billy/user_service/internal/models"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func (repo *PostgresUserRepository) Create(user *models.User) (int, error) {
	query := `INSERT INTO users_db (
        name, 
        login, 
        payment_acc, 
        conn_type, 
        balance, 
        write_off_date, 
        active, 
        plan_title, 
        plan_series, 
        plan_subs_fee, 
        last_payment_sum, 
        last_payment_date
    ) VALUES (
        $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
    ) RETURNING id;`
	var id int
	err := repo.db.QueryRow(query,
		user.Name, user.Login, user.PaymentAcc, user.ConnType, user.Balance, user.WriteOffDate, user.Active, user.PlanTitle, user.PlanSeries, user.PlanSubsFee, user.LastPaymentSum, user.LastPaymentDate).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (repo *PostgresUserRepository) FindByID(id string) (*models.User, error) {
	query := `
		SELECT id, name, login, payment_acc, conn_type, balance, write_off_date, active,
			   plan_title, plan_series, plan_subs_fee, last_payment_sum, last_payment_date,
			   created_at, updated_at
		FROM users_db
		WHERE id = $1;
	`

	var user models.User
	err := repo.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Login,
		&user.PaymentAcc,
		&user.ConnType,
		&user.Balance,
		&user.WriteOffDate,
		&user.Active,
		&user.PlanTitle,
		&user.PlanSeries,
		&user.PlanSubsFee,
		&user.LastPaymentSum,
		&user.LastPaymentDate,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return &user, nil
}

func (repo *PostgresUserRepository) FindAll() ([]models.User, error) {
	query := `
		SELECT id, name, login, payment_acc, conn_type, balance, write_off_date, active,
			   plan_title, plan_series, plan_subs_fee, last_payment_sum, last_payment_date,
			   created_at, updated_at
		FROM users_db;
	`
	users := []models.User{}
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID,
			&user.Name,
			&user.Login,
			&user.PaymentAcc,
			&user.ConnType,
			&user.Balance,
			&user.WriteOffDate,
			&user.Active,
			&user.PlanTitle,
			&user.PlanSeries,
			&user.PlanSubsFee,
			&user.LastPaymentSum,
			&user.LastPaymentDate,
			&user.CreatedAt,
			&user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (repo *PostgresUserRepository) Update(id string, user *models.User) error {
	query := `
		UPDATE users_db
		SET name = $1,
		    login = $2,
		    conn_type = $3,
		    balance = $4,
		    write_off_date = $5,
		    active = $6,
		    plan_title = $7,
		    plan_series = $8,
		    plan_subs_fee = $9,
		    last_payment_sum = $10,
		    last_payment_date = $11,
		    updated_at = $12
		WHERE id = $13;
	`

	_, err := repo.db.Exec(query,
		user.Name,
		user.Login,
		user.ConnType,
		user.Balance,
		user.WriteOffDate,
		user.Active,
		user.PlanTitle,
		user.PlanSeries,
		user.PlanSubsFee,
		user.LastPaymentSum,
		user.LastPaymentDate,
		time.Now(),
		id,
	)

	if err != nil {
		return fmt.Errorf("failed to update user with id %s: %w", id, err)
	}

	return nil
}

func (repo *PostgresUserRepository) Delete(id string) error {
	query := `
	DELETE FROM users_db
	WHERE id = $1;
`

	result, err := repo.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user with id %s: %w", id, err)
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to determine affected rows: %w", err)
	}

	if affectedRows == 0 {
		return fmt.Errorf("no user found with id %s", id)
	}

	return nil
}
