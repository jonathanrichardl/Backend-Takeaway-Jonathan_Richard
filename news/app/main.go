package main

import (
	"log"
	config "news/config"
	databaseinstance "news/pkg/database"
	"news/pkg/handlers"
	"news/pkg/router"
)

func main() {
	// SQL credentials are supplied in config/database.json
	log.Println("Reading Database Configuration")
	databaseConfig, err := config.LoadDatabaseConfiguration()
	if err != nil {
		log.Printf("Error setting database : %s\n", err.Error())
		return
	}

	//initializing db and router
	log.Println("Initializing Program")
	Database, err := databaseinstance.NewDatabase(databaseConfig.DatabaseManagementSystem,
		databaseConfig.Username, databaseConfig.Password, databaseConfig.Address,
		databaseConfig.DatabaseName)

	if err != nil {
		log.Printf("Error received : %s\n", err.Error())
		return
	}
	Router := router.NewRouterInstance()
	handlers := handlers.NewHttpHandlers(Database, Router)
	handlers.RegisterAllHandlers()
}
