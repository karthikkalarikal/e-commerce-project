package http

import (
	"github.com/gin-gonic/gin"
	"github.com/karthikkalarikal/ecommerce-project/pkg/api/handler"
	"github.com/karthikkalarikal/ecommerce-project/pkg/routes"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(UserHandler *handler.UserHandler) *ServerHTTP {
	engine := gin.New()

	engine.Use(gin.Logger())

	routes.UserRoutes(engine.Group("/users"), UserHandler)
	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":8080")
}
