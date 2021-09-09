package main

import (
	"log"
	databaseinstance "news/pkg/database"
	"news/pkg/handlers"
	"news/pkg/router"
)

func main() {
	// SQL credentials are supplied in config/database.json
	log.Println("Reading Database Configuration")
	/*databaseConfig, err := config.LoadDatabaseConfiguration()
	if err != nil {
		log.Printf("Error setting database : %s\n", err.Error())
		return
	}
	*/

	//initializing db and router
	log.Println("Initializing Program")
	Database, err := databaseinstance.NewDatabase("mysql",
		"root", "123jonathan123100300!!!", "localhost:3306",
		"testers")

	if err != nil {
		log.Printf("Error received : %s\n", err.Error())
		return
	}
	Router := router.NewRouterInstance()
	handlers := handlers.NewHttpHandlers(Database, Router)
	handlers.RegisterAllHandlers()
	Router.Start()
}
