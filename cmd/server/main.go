/*
main.go
Author : Naveenraj O M
Description: This is the entry point for the M Commerce application. It initializes the configuration, MongoDB connection, and the HTTP server.
*/
package main

import (
	"fmt"
	"log"
	"mcommerce/config"
	"mcommerce/internal/constants"
	"mcommerce/internal/database"
	"mcommerce/internal/firebase"
	"mcommerce/internal/router"
)

func main() {
	// Initialize application configuration from environment variables
	config.LoadConfig()

	// If the environment is dev making extra timeout
	if config.Config.Environment == "dev" {
		constants.DatabaseTimeout = 1000
	}

	// Initialize MongoDB establishes a connection to the MongoDB database using the connection URI
	database.InitializeMongoDB()
	defer database.Disconnect()

	// Initialize MongoDB collections by using the MongoDB Database
	database.InitializeCollections()

	// Initialize Firebase
	firebase.InitFirebase()

	router := router.GetRouter()
	if err := router.Listen(fmt.Sprintf(":%d", config.Config.Port)); err != nil {
		log.Println("Failed to start server")
	}
}
