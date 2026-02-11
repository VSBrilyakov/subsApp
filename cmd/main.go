package main

import (
	"log"

	"githhub.com/VSBrilyakov/test-app/configs"
	"githhub.com/VSBrilyakov/test-app/internal"
	"githhub.com/VSBrilyakov/test-app/internal/handler"
	"githhub.com/VSBrilyakov/test-app/internal/repository"
	"githhub.com/VSBrilyakov/test-app/internal/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		PrettyPrint:     true,
	})
	logrus.SetLevel(logrus.InfoLevel)

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
	logrus.Info("config loaded")

	db, err := repository.NewPostgresDB(&config.Postgres)
	if err != nil {
		log.Fatalf("Postgres connection error: %s", err.Error())
	}
	logrus.Info("postgres connection established")

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := new(internal.Server)
	if err := srv.Run(config.Server, handlers.InitRoutes()); err != nil {
		log.Fatalf("server run failed: %s", err.Error())
	}
}
