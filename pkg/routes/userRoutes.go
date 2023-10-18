package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/karthikkalarikal/ecommerce-project/pkg/api/handler"
	"github.com/karthikkalarikal/ecommerce-project/pkg/api/middlewar"
)

func UserRoutes(engine *gin.RouterGroup, userHandler *handler.UserHandler, otpHandler *handler.OtpHandler, productHandler *handler.ProductHandler, cartHandler *handler.CartHandler, orderHandler *handler.OrderHandler, payment *handler.PaymentHandler) {
	engine.POST("/signup", userHandler.UserSignUp)
	engine.POST("/login", userHandler.LoginHandler)

	engine.POST("/otplogin", otpHandler.SendOTP)
	engine.POST("/verifyotp", otpHandler.VerifyOTP)

	engine.GET("/viewproducts", productHandler.ListProducts)
	engine.GET("/viewbycategories", productHandler.ListByCategoreis)
	engine.GET("/payment", payment.MakePaymentRazorpay)
	engine.GET("/verifypayment", payment.VerifyPayment)

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
			addressmanagement.PATCH("/select", userHandler.SelectAddress)
		}
		usermanagement := engine.Group("/user")
		{
			usermanagement.GET("/viewdetails", userHandler.ViewUser)
			usermanagement.GET("/addresses", userHandler.GetAddress)
			usermanagement.PUT("/edit", userHandler.EditUserDetails)
			usermanagement.POST("/changepassword", userHandler.ChangePassword)
		}
		ordermanagement := engine.Group("/order")
		{
			ordermanagement.POST("/add", orderHandler.AddToOrder)
			ordermanagement.GET("/view", orderHandler.ViewOrder)
			// ordermanagement.POST("/verifypayment", payment.VerifyPayment)
		}
	}

}
