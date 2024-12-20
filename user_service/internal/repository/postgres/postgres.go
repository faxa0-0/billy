package postgres

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}
