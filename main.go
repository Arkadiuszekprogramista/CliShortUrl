package main

import (
	"app/cmd/shorterCmd"
	"app/pkg/config"
	"app/pkg/shorter"
	"log"
)

var Redis shorter.Redis

func main() {


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
