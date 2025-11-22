package main

import (
	"log"

	"blog/config"
	"blog/database"
	"blog/routes"
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

	// Setup routes and start server
	ginServer := routes.SetupRoutes()
	ginServer.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
