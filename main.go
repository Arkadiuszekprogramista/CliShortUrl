package main

import (
	"app/pkg/shorter"
	"app/cmd/shorterCmd"
	"log"

)


func main() {

	configuration, err := ConfigFromFile("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	service, err := shorter.NewPool(configuration.Redis.Host, configuration.Redis.Port)
	if err != nil {
		log.Fatal(err)
	}
	defer service.Close()
	shorterCmd.Execute()

}
