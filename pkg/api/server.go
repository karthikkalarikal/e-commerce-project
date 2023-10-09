package http

import (
	"github.com/gin-gonic/gin"
	_ "github.com/karthikkalarikal/ecommerce-project/cmd/api/docs"
	"github.com/karthikkalarikal/ecommerce-project/pkg/api/handler"
	"github.com/karthikkalarikal/ecommerce-project/pkg/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler, otpHandler *handler.OtpHandler, productHandler *handler.ProductHandler, adminHandler *handler.AdminHandler, cartHandler *handler.CartHandler) *ServerHTTP {
	engine := gin.New()

	engine.Use(gin.Logger())

	//swagger docs
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.UserRoutes(engine.Group("/users"), userHandler, otpHandler, productHandler, cartHandler)
	routes.AdminRoutes(engine.Group("/admin"), adminHandler, productHandler)
	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":8080")
}
