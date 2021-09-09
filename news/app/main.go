package main

import (
	"log"
	"news/config"
	"news/pkg/database"
	"news/pkg/handlers"
	logger "news/pkg/logger"
	"news/pkg/router"
)

func main() {
	Logger := logger.NewLogger("log.txt")
	Logger.InfoLogger.Println("Reading database configuration")

	databaseConfig, err := config.LoadDatabaseConfiguration()
	if err != nil {
		log.Printf("Error setting database : %s\n", err.Error())
		return
	}

	//initializing db and router
	Logger.InfoLogger.Println("Initializing Program")
	Database, err := database.NewDatabase("mysql",
		databaseConfig.Username, databaseConfig.Password, databaseConfig.Address,
		databaseConfig.DatabaseName)

	if err != nil {
		log.Printf("Error received : %s\n", err.Error())
		return
	}
	Router := router.NewRouterInstance()
	handlers := handlers.NewHttpHandlers(Database, Router, Logger)
	handlers.RegisterAllHandlers()
	Router.Start()
}
