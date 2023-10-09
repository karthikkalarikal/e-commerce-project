package main

import (
	"log"

	"github.com/karthikkalarikal/ecommerce-project/pkg/config"
	"github.com/karthikkalarikal/ecommerce-project/pkg/di"
)

func main() {

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	server, err := di.InitializeAPI(config)
	if err != nil {
		log.Fatal(err)
	} else {
		server.Start()
	}
}
