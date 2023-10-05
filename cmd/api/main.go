package main

import (
	"log"

	"github.com/karthikkalarikal/ecommerce-project/cmd/api/docs"
	"github.com/karthikkalarikal/ecommerce-project/pkg/config"
	"github.com/karthikkalarikal/ecommerce-project/pkg/di"
)

func main() {
	docs.SwaggerInfo.Title = "TechDeck"
	docs.SwaggerInfo.Description = "Tech Store -- E-commerce"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https", "ssl"}
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
