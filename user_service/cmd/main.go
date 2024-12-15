package main

import (
	"log"

	"github.com/faxa0-0/billy/plan_service/internal/config"
	"github.com/faxa0-0/billy/plan_service/internal/repository"
	"github.com/faxa0-0/billy/plan_service/internal/repository/sqlite"
)

func main() {
	//TODO: Config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	//TODO: Database
	var repo repository.Repository
	repo, err = sqlite.NewSQLiteUserRepository(cfg.DBPath)
	if err != nil {
		log.Fatal("failed to initialize repository")
	}
	//TODO: Service
	_ = repo
	//TODO: API
	//TODO: Start APP
}
