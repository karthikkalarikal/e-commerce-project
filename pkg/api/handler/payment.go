package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	pay "github.com/karthikkalarikal/ecommerce-project/pkg/usecase/interfaces"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/response"
)

type PaymentHandler struct {
	payment pay.PaymentUseCase
}

func NewPaymentHandler(useCase pay.PaymentUseCase) *PaymentHandler {
	return &PaymentHandler{
		payment: useCase,
	}
}

// @Summary Delete product
// @Description Delete product by id
// @Tags Order Management
// @Accept json
// @Produce json
// @Param user_id query int true "user id"
// @Param order_id query int true "order_id"
// @Security BearerTokenAuth
// @Success 200 {object} response.Response "invoice with details of order and user "
// @Failure 400 {object} response.Response  "Bad request"
// @Router /users/order/payment [post]
func (handler *PaymentHandler) MakePaymentRazorpay(c *gin.Context) {
	userId := c.Query("user_id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {

		errRes := response.ClientResponse(http.StatusBadRequest, "error", nil, errors.New("error in converting string to int userid"+err.Error()))
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	orderId := c.Query("order_id")
	orderIdInt, err := strconv.Atoi(orderId)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error", nil, errors.New("error in converting string to int orderid"+err.Error()))
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	body, razorId, err := handler.payment.MakePaymentRazorpay(orderIdInt, userIdInt)
	fmt.Println("body in handler", body)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	fmt.Println("body next", body.Amount, razorId, userId, body.OrderId, body.Name, body.Amount)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"final_price": body.Amount * 100,
		"razor_id":    razorId,
		"user_id":     userId,
		"order_id":    body.OrderId,
		"user_name":   body.Name,
		"total":       int(body.Amount),
	})
}
