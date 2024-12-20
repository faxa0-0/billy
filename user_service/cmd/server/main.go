package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/faxa0-0/billy/user_service/internal/config"
	"github.com/faxa0-0/billy/user_service/internal/handler"
	"github.com/faxa0-0/billy/user_service/internal/middleware"
	"github.com/faxa0-0/billy/user_service/internal/repository/postgres"
	"github.com/faxa0-0/billy/user_service/internal/service"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	cfg, err := config.LoadYAML()
	if err != nil {
		log.Fatalf("unable to load config because %s", err)
	}

	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.Database.Username, cfg.Database.Password, cfg.Database.Host,
		cfg.Database.Port, cfg.Database.DBName, cfg.Database.SSLMode)

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database because %s", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatalf("Cannot ping database because %s", err)
	}

	log.Println("Successfully connected to database and pinged it")

	repo := postgres.NewPostgresUserRepository(db)
	service := service.NewUserService(repo)
	handler := handler.NewUserHandler(service)

	mux := http.NewServeMux()

	mux.Handle("/users", middleware.AuthMiddleware(http.HandlerFunc(handler.HandleUserRequests)))
	mux.Handle("/users/", middleware.AuthMiddleware(http.HandlerFunc(handler.HandleUserRequests)))

	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	log.Println("Starting server on", addr)
	err = http.ListenAndServe(addr, mux)
	if err != nil {
		log.Fatalf("server failed to start: %v", err)
	}

}
