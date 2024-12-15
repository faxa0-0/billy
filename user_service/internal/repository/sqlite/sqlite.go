package sqlite

import (
	"github.com/faxa0-0/billy/plan_service/internal/repository"
)

type SQLiteUserRepository struct {
	//db *sql.DB
}

func NewSQLiteUserRepository(path string) (repository.Repository, error) {
	return &SQLiteUserRepository{}, nil
}
