package main

import (
	"log"
	"net/http"

	"mockapi/internal/config"
	"mockapi/internal/database"
	"mockapi/internal/mockmock"
	"mockapi/internal/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	if err := database.RunMigrations(cfg.DBURL); err != nil {
		log.Fatalf("run migrations: %v", err)
	}

	pool, err := database.Connect(cfg.DBURL)
	if err != nil {
		log.Fatalf("connect db: %v", err)
	}
	defer pool.Close()

	repo := mockmock.NewPostgres(pool)
	svc := mockmock.NewService(repo)
	handler := mockmock.NewHandler(svc)
	router := server.NewRouter(handler)

	log.Printf("server listening on :%s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
