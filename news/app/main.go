package main

import (
	"fmt"
	"log"
	"news/config"
	"news/pkg/database"
	"news/pkg/handlers"
	logger "news/pkg/logger"
	"news/pkg/redis"
	"news/pkg/router"
)

func main() {
	Logger := logger.NewLogger("log.txt")
	fmt.Println("Starting")
	Logger.InfoLogger.Println("Reading database configuration")
	/*
		databaseConfig, err := config.LoadDatabaseConfiguration()
		if err != nil {
			log.Printf("Error setting database : %s\n", err.Error())
			return
		}
	*/
	Logger.InfoLogger.Println("Reading server configuration")
	serverConfig, err := config.LoadServerConfiguration()
	if err != nil {
		log.Printf("Error setting database : %s\n", err.Error())
		return
	}
	//initializing db and router
	Logger.InfoLogger.Println("Initializing Program")
	/*
		Database, err := database.NewDatabase("mysql",
			databaseConfig.Username, databaseConfig.Password, databaseConfig.Address,
			databaseConfig.DatabaseName)
	*/
	database, err := database.NewDatabase("mysql",
		"newuser", "123Jonathan123100300!!!", "localhost:3306",
		"testers")
	if err != nil {
		log.Printf("Error received : %s\n", err.Error())
		return
	}
	redis := redis.NewRedisClient(serverConfig.Redis.Host, serverConfig.Redis.Port,
		serverConfig.Redis.Password, serverConfig.Redis.DatabaseNo,
		serverConfig.Redis.Expiration)
	router := router.NewRouterInstance()
	handlers := handlers.NewHttpHandlers(database, router, Logger, redis)
	handlers.RegisterAllHandlers()
	router.Start(fmt.Sprintf("%s:%d", serverConfig.Server.Host, serverConfig.Server.Port))
}
