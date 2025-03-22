package main

import (
	"explorekg/internal/app"
	"log"
)

func main() {
	router := app.App()
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
