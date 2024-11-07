/*
mongodb.go
Author : Naveenraj O M
Description : This file manages MongoDB connections and collections for the application. It includes functions to initialize and disconnect the MongoDB client, and provides references to collections used in the application.
*/
package database

import (
	"context"
	"log"
	"mcommerce/config"
	"mcommerce/internal/constants"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoClient *mongo.Client
	DB          *mongo.Database
	mongoOnce   sync.Once
)

var Collections DatabaseCollections

type DatabaseCollections struct {
	Users *mongo.Collection
}

func InitializeMongoDB() {
	mongoOnce.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), constants.DatabaseTimeout*time.Second)
		defer cancel()

		client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Config.DBUri))
		if err != nil {
			log.Fatalf("Failed to connect to MongoDB: %v", err)
		}

		// Ping the database
		err = client.Ping(ctx, nil)
		if err != nil {
			log.Fatalf("Failed to ping MongoDB: %v", err)
		}

		MongoClient = client
		DB = client.Database(config.Config.DBName)
	})

	log.Println("MongoDB collections initialized successfully")
}

func Disconnect() {
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), constants.DatabaseTimeout*time.Second)
	defer cancel()

	if MongoClient != nil {
		if err := MongoClient.Disconnect(ctx); err != nil {
			log.Printf("Error disconnecting from database: %v", err)
		}
	}

	log.Println("Successfully disconnected from MongoDB")
}

func InitializeCollections() {
	// add the all collections here
	Collections = DatabaseCollections{
		Users: DB.Collection("users"),
	}
}
