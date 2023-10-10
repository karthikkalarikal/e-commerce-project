package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/karthikkalarikal/ecommerce-project/pkg/api/handler"
	"github.com/karthikkalarikal/ecommerce-project/pkg/api/middlewar"
)

func UserRoutes(engine *gin.RouterGroup, userHandler *handler.UserHandler, otpHandler *handler.OtpHandler, productHandler *handler.ProductHandler, cartHandler *handler.CartHandler) {
	engine.POST("/signup", userHandler.UserSignUp)
	engine.POST("/login", userHandler.LoginHandler)

	engine.POST("/otplogin", otpHandler.SendOTP)
	engine.POST("/verifyotp", otpHandler.VerifyOTP)

	engine.GET("/viewproducts", productHandler.ListProducts)
	engine.GET("/viewbycategories", productHandler.ListByCategoreis)

	engine.Use(middlewar.UserMiddleware)
	{
		cartmanagement := engine.Group("/carts")
		{
			cartmanagement.POST("/addtocart", cartHandler.AddToCart)
			cartmanagement.GET("/viewcart", cartHandler.CartItemListing)
			cartmanagement.PATCH("/quantity", cartHandler.CartItemQuatityModification)
			cartmanagement.DELETE("/delete", cartHandler.CartItemDeletion)
		}
		addressmanagement := engine.Group("/user")
		{
			addressmanagement.POST("/address", userHandler.UserAddress)
		}
	}

}
