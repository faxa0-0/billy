package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(path string) (*SQLiteRepository, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS plans (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			series TEXT,
			corp INTEGER,
			subscription_fee INTEGER,
			conn_type TEXT,
			speed TEXT,
			tasix_speed TEXT,
			limit_mb INTEGER,
			additional_info TEXT,
			active BOOLEAN,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return nil, err
	}

	return &SQLiteRepository{db: db}, nil
}
func (repo *SQLiteRepository) Close() error {
	return repo.db.Close()
}
