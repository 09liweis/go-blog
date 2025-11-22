package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

// ConnectToMongoDB connects to MongoDB database
func ConnectToMongoDB(uri string) error {
	log.Printf("Attempting to connect to MongoDB...")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Printf("MongoDB connection error: %v", err)
		return err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Printf("MongoDB ping error: %v", err)
		return err
	}

	log.Printf("Successfully connected to MongoDB!")
	mongoClient = client
	return err
}

// DisconnectFromMongoDB disconnects from MongoDB
func DisconnectFromMongoDB() error {
	if mongoClient != nil {
		return mongoClient.Disconnect(context.TODO())
	}
	return nil
}

// GetMongoClient returns the MongoDB client instance
func GetMongoClient() *mongo.Client {
	return mongoClient
}