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
			usermanagement.PATCH("/block/:user_id", adminHandler.BlockUser)
			usermanagement.GET("/searchuser", adminHandler.FindUser)
			usermanagement.DELETE("/deleteuser/:user_id", adminHandler.DeleteUser)
		}
		productmanagement := engine.Group("/product")
		{
			productmanagement.POST("/addproduct", productHandler.AddProduct)
			productmanagement.POST("/editproduct", adminHandler.EditProduct)
			productmanagement.DELETE("/deleteproduct/:product_id", productHandler.DeleteProduct)
			productmanagement.POST("/addcategory", productHandler.AddCategory)
			productmanagement.PUT("/updatecategory/:id", productHandler.UpdateCategory)
			productmanagement.DELETE("/deletecategory/:id", productHandler.DeleteCategory)
		}
	}

}
