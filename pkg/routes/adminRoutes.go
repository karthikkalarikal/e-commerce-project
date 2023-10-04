package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/karthikkalarikal/ecommerce-project/pkg/api/handler"
)

func AdminRoutes(engine *gin.RouterGroup, adminHandler *handler.AdminHandler) {
	engine.POST("/userlist", adminHandler.UserList)
}
