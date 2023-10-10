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

	engine.Use(middlewar.AdminMiddleware)
	{
		usermanagement := engine.Group("/user")
		{
			usermanagement.POST("/addtocart", cartHandler.AddToCart)
		}
	}

}
