package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadConfig loads environment variables and returns MongoDB URI
func LoadConfig() string {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Reload URI after loading .env file
	uri := os.Getenv("MONGODB_URL")
	log.Printf("MONGODB_URL: %s", uri)
	log.Printf("URI length: %d", len(uri))

	if uri == "" {
		log.Fatal("You must set your 'MONGO_URL' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	return uri
}