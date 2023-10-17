package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/karthikkalarikal/ecommerce-project/pkg/usecase/interfaces"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/response"
)

type OrderHandler struct {
	orderUseCase interfaces.OrderUseCase
}

func NewOrderHandler(usecase interfaces.OrderUseCase) *OrderHandler {
	return &OrderHandler{
		orderUseCase: usecase,
	}
}

// -------------------------------------- insert order handler ------------------------------------------------ \\

// @Summary Add To Order
// @Description Add cart to the order using user id and cart id
// @Tags Order Management
// @Accept json
// @Produce json
// @Param cart_id query int true "cart id"
// @Param user_id query int true "user_id"
// @Security BearerTokenAuth
// @Success 200 {object} response.Response "success"
// @Failure 500 {object} response.Response{} "fail"
// @Router /users/order/add [post]
func (handler *OrderHandler) AddToOrder(ctx *gin.Context) {

	userId := ctx.Query("user_id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in user id", nil, err.Error())
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}

	cartId := ctx.Query("cart_id")
	cartIdInt, err := strconv.Atoi(cartId)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in cart id", nil, err.Error())
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}

	resOrder, err := handler.orderUseCase.AddToOrder(userIdInt, cartIdInt)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in ordering", nil, err.Error())
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "succesfully added order", resOrder, nil)
	ctx.JSON(http.StatusOK, successRes)

}

// ---------------------------------------- change status of order -------------------------------------- \\

// @Summary Add To Order
// @Description Add cart to the order using user id and cart id
// @Tags Order Mangement
// @Accept json
// @Produce json
// @Param order_id query int true "order_id"
// @Security BearerTokenAuth
// @Success 200 {object} response.Response "success"
// @Failure 500 {object} response.Response{} "fail"
// @Router /users/order/cancel [delete]
// func (handler *OrderHandler) CancelOrder(ctx *gin.Context) {
// 	order_id := ctx.Query("order_id")
// 	ordrId, err := strconv.Atoi(order_id)
// 	if err != nil {
// 		errRes := response.ClientResponse(http.StatusBadGateway, "error in reading the order id", nil, err.Error())
// 		ctx.JSON(http.StatusBadRequest, errRes)
// 		return
// 	}

// }
