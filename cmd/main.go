package main

import (
	_ "explorekg/docs"
	"explorekg/internal/app"
	"log"
)

// @title Explore-Base API documentation
// @version 1.0.1
// @host localhost:8080
// @BasePath
func main() {
	router := app.App()
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
