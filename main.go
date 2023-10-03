package main

import (
	"app/cmd/shorterCmd"
	"app/pkg/config"
	"app/pkg/shorter"
	"log"
)

func main() {

	var app config.AppConfig

	// app config
	app.UseMySite = false
	app.MySite = "https://www.wp.pl/"

	shorterCmd.NewConfig(&app)

	configuration, err := config.ConfigFromFile("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	db, err := shorter.NewPool(configuration.Redis.Host, configuration.Redis.Port)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Pool.Close()

	shorterCmd.Execute()

}
