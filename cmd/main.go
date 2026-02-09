package main

import (
	"log"

	"githhub.com/VSBrilyakov/test-app/internal"
	"githhub.com/VSBrilyakov/test-app/internal/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(internal.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("server run failed: %s", err.Error())
	}
}
