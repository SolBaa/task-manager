package main

import (
	"log"
	"net/http"

	"github.com/SolBaa/task-manager/config"
	"github.com/SolBaa/task-manager/internal/db"
	"github.com/SolBaa/task-manager/internal/routes"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {

	// Setup logging
	// logger.Init() //TODO
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Load configuration
	cfg := config.LoadConfig()

	// Initialize services, repositories, handlers, etc.
	db, err := db.Initialize(*&cfg.DBConfig)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	defer db.Close()
	r := chi.NewRouter()
	// Setup routes
	router := routes.SetupRouter(r, db)

	// Start server
	log.Println("Starting server on port", cfg.ServerPort)
	if err := http.ListenAndServe(":"+cfg.ServerPort, router); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}
