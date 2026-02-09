package main

import (
	"log"

	"githhub.com/VSBrilyakov/test-app/configs"
	"githhub.com/VSBrilyakov/test-app/internal"
	"githhub.com/VSBrilyakov/test-app/internal/handler"
	"githhub.com/VSBrilyakov/test-app/internal/repository"
	"githhub.com/VSBrilyakov/test-app/internal/service"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}
}

func main() {
	config, err := configs.NewConfig()
	if err != nil {
		log.Fatalf("Config reading error: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(&config.Postgres)
	if err != nil {
		log.Fatalf("Postgres connection error: %s", err.Error())
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := new(internal.Server)
	if err := srv.Run(config.Server, handlers.InitRoutes()); err != nil {
		log.Fatalf("server run failed: %s", err.Error())
	}
}
