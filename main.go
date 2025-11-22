package main

import (
	"log"

	"blog/config"
	"blog/database"
	"blog/routes"
	"blog/handlers"
)

//Todo: add favicon

func main() {
	// Load configuration
	uri := config.LoadConfig()

	// Connect to MongoDB
	if err := database.ConnectToMongoDB(uri); err != nil {
		log.Fatal("Could not connect to MongoDB")
	}
	defer func() {
		if err := database.DisconnectFromMongoDB(); err != nil {
			log.Printf("Error disconnecting from MongoDB: %v", err)
		}
	}()

	// Seed sample data if database is empty
	if err := handlers.SeedBlogs(); err != nil {
		log.Printf("Warning: Could not seed sample data: %v", err)
	}

	// Setup routes and start server
	ginServer := routes.SetupRoutes()
	ginServer.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
