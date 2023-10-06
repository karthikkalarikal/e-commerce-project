package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/karthikkalarikal/ecommerce-project/pkg/api/handler"
	"github.com/karthikkalarikal/ecommerce-project/pkg/api/middlewar"
)

func AdminRoutes(engine *gin.RouterGroup, adminHandler *handler.AdminHandler) {

	engine.Use(middlewar.AdminMiddleware)
	{
		usermanagement := engine.Group("/users")
		{
			usermanagement.POST("/userlist", adminHandler.UserList)
			usermanagement.POST("/block", adminHandler.BlockUser)
		}
	}

}
