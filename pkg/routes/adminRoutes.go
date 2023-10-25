package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/karthikkalarikal/ecommerce-project/pkg/api/handler"
	"github.com/karthikkalarikal/ecommerce-project/pkg/api/middlewar"
)

func AdminRoutes(engine *gin.RouterGroup, adminHandler *handler.AdminHandler, productHandler *handler.ProductHandler, couponHandler *handler.CouponHandler) {

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
			productmanagement.PUT("/editproduct/:product_id", productHandler.EditProduct)
			productmanagement.DELETE("/deleteproduct/:product_id", productHandler.DeleteProduct)
			productmanagement.POST("/addcategory", productHandler.AddCategory)
			productmanagement.PUT("/updatecategory/:id", productHandler.UpdateCategory)
			productmanagement.DELETE("/deletecategory/:category_id", productHandler.DeleteCategory)
		}
		couponmanagement := engine.Group("/coupon")
		{
			couponmanagement.POST("/addcoupon", couponHandler.AddCoupon)
			couponmanagement.GET("/view", couponHandler.ViewCoupon)
			couponmanagement.PATCH("/expire", couponHandler.ExpireCoupon)
		}
		ordermanagement := engine.Group("/dashboard")
		{
			ordermanagement.GET("/totalsales", adminHandler.GetTotalAmount)
			ordermanagement.GET("/salesbydate", adminHandler.GetSalesDetailsByDate)
		}
	}

}
