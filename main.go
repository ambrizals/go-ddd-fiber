package main

import (
	"github.com/ambrizals/go-ddd-fiber/config"
	"github.com/ambrizals/go-ddd-fiber/infra/db"
	"github.com/ambrizals/go-ddd-fiber/routes"
	"log"
)

func main() {
	// Initialize environment variables
	_, envErr := config.LoadEnv()
	if envErr != nil {
		log.Fatalf("Failed to load environment variables: %v", envErr)
	}
	log.Println("Environment variables loaded")

	// Initialize database client
	_, err := db.GetClient()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Initialize and start the server with dependencies
	app := routes.New()
	log.Println("Server starting on :3000...")
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
