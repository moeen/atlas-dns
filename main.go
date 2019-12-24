package main

import (
	"github.com/moeen/atlas-dns/config"
	"github.com/moeen/atlas-dns/router"
	"log"
)

func init() {
	// Loading the config from environment variables and fatal if there was an error
	if err := config.LoadConfig(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Running the http server with the port defined in the config
	log.Fatal(router.Start(config.GetPort()))
}
