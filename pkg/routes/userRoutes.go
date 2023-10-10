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
		usermanagement := engine.Group("/carts")
		{
			usermanagement.POST("/addtocart", cartHandler.AddToCart)
			usermanagement.GET("/viewcart", cartHandler.CartItemListing)
			usermanagement.PATCH("/quantity", cartHandler.CartItemQuatityModification)
		}
	}

}
