package main

import (
	"log"

	"github.com/faxa0-0/billy/plan_service/internal/api"
	"github.com/faxa0-0/billy/plan_service/internal/config"
	"github.com/faxa0-0/billy/plan_service/internal/handler"
	"github.com/faxa0-0/billy/plan_service/internal/repository/sqlite"
	"github.com/faxa0-0/billy/plan_service/internal/service"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	repo, err := sqlite.NewSQLiteRepository(cfg.DBPath)
	if err != nil {
		log.Fatal("failed to initialize repository")
	}
	defer repo.Close()

	planService := service.NewPlanService(repo)

	planHandler := handler.NewPlanHandler(planService)

	api := api.NewApi(*planHandler, cfg.Address)

	api.Run()
}
