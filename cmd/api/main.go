package main

import (
	"log"

	"github.com/karthikkalarikal/ecommerce-project/cmd/api/docs"
	"github.com/karthikkalarikal/ecommerce-project/pkg/config"
	"github.com/karthikkalarikal/ecommerce-project/pkg/di"
)

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @securityDefinitions.basic  BasicAuth

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
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
