package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/VSBrilyakov/subsApp/configs"
	"github.com/VSBrilyakov/subsApp/docs"
	"github.com/VSBrilyakov/subsApp/internal"
	"github.com/VSBrilyakov/subsApp/internal/handler"
	"github.com/VSBrilyakov/subsApp/internal/repository"
	"github.com/VSBrilyakov/subsApp/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var config *configs.Config

func init() {
	var err error

	err = godotenv.Load()
	if err != nil {
		logrus.Fatal("invalid .env file")
	}

	config, err = configs.NewConfig()
	if err != nil {
		logrus.Fatalf("config reading error: %s", err.Error())
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:               true,
		DisableColors:             false,
		EnvironmentOverrideColors: true,
		FullTimestamp:             true,
	})
	var logLvl logrus.Level
	if logLvl, err = logrus.ParseLevel(config.LogLevel); err != nil {
		logrus.Fatalf("invalid log level: %s", err.Error())
	}
	logrus.SetLevel(logLvl)
	logrus.Info("config loaded")

	docs.SwaggerInfo.Host = "localhost:" + strconv.Itoa(config.Server.Port)
	docs.SwaggerInfo.BasePath = "/"

	gin.SetMode(gin.ReleaseMode)
}

// @title Subs App API
// @version 1.0
// @description API Server for SubsApp Application
func main() {
	db, err := repository.NewPostgresDB(&config.Postgres)
	if err != nil {
		logrus.Fatalf("postgres connection error: %s", err.Error())
	}
	//close connection to db after exiting main()
	defer func() {
		err = db.Close()
		if err := db.Close(); err != nil {
			logrus.Errorf("error occured on db connection close: %s", err.Error())
		}
	}()
	logrus.Info("postgres connection established")

	err = repository.DoMigrates(db)
	if err != nil {
		logrus.Fatalf("migrations applying error: %s", err.Error())
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := new(internal.Server)
	go func() {
		if err := srv.Run(config.Server, handlers.InitRoutes()); !errors.Is(err, http.ErrServerClosed) {
			logrus.Fatalf("server run failed: %s", err.Error())
		}
	}()
	logrus.Info("SubsApp started")

	//graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Info("SubsApp shutting down")

	if err := srv.Stop(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
}
