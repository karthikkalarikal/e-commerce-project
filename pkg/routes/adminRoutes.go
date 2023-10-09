package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/karthikkalarikal/ecommerce-project/pkg/api/handler"
	"github.com/karthikkalarikal/ecommerce-project/pkg/api/middlewar"
)

func AdminRoutes(engine *gin.RouterGroup, adminHandler *handler.AdminHandler, productHandler *handler.ProductHandler) {

	engine.Use(middlewar.AdminMiddleware)
	{
		usermanagement := engine.Group("/users")
		{
			usermanagement.GET("/userlist", adminHandler.UserList)
			usermanagement.POST("/block", adminHandler.BlockUser)
			usermanagement.GET("/searchuser", adminHandler.FindUser)
			usermanagement.POST("/deleteuser", adminHandler.DeleteUser)
		}
		productmanagement := engine.Group("/product")
		{
			productmanagement.POST("/addproduct", adminHandler.AddProduct)
			productmanagement.POST("/editproduct", adminHandler.EditProduct)
			productmanagement.POST("/deleteproduct/:id", adminHandler.DeleteProduct)
			productmanagement.POST("/addcategory", adminHandler.AddCategory)
			productmanagement.PUT("/updatecategory/:id", productHandler.UpdateCategory)
			productmanagement.DELETE("/deletecategory/:id", productHandler.DeleteCategory)
		}
	}

}
