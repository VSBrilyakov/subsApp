package main

import (
	"log"

	"githhub.com/VSBrilyakov/test-app/internal"
)

func main() {
	srv := new(internal.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("server run failed: %s", err.Error())
	}
}
