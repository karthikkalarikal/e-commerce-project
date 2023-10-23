package handler

import (
	"errors"
	"fmt"
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
// @Tags Order Management
// @Accept json
// @Produce json
// @Param order_id query int true "order_id"
// @Security BearerTokenAuth
// @Success 200 {object} response.Response "success"
// @Failure 500 {object} response.Response{} "fail"
// @Router /users/order/view [get]
func (handler *OrderHandler) ViewOrder(ctx *gin.Context) {
	order_id := ctx.Query("order_id")
	ordrId, err := strconv.Atoi(order_id)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadGateway, "error in reading the order id", nil, err.Error())
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}

	body, err := handler.orderUseCase.ViewOrder(ordrId)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadGateway, "error in getting order", nil, err.Error())
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "the request was succesful", body, nil)
	ctx.JSON(http.StatusOK, successRes)

}

// --------------------------------------------------- cancel order ------------------------------------------- \\

// @Summary Cancel Order
// @Description Cancel Order By Order Id
// @Tags Order Management
// @Produce json
// @Param order_id query int true "order_id"
// @Security BearerTokenAuth
// @Success 200 {object} response.Response "success"
// @Failure 500 {object} response.Response{} "fail"
// @Router /users/order/cancel [delete]
func (handler *OrderHandler) CancelOrder(ctx *gin.Context) {
	orderId := ctx.Query("order_id")
	orderIdInt, err := strconv.Atoi(orderId)
	fmt.Println("order id", orderIdInt)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadGateway, "error in reading the order id", nil, err.Error())
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}

	body1, body2, err := handler.orderUseCase.CancelOrder(orderIdInt)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadGateway, "error in canceling ", nil, err.Error())
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}
	var body []interface{}
	body = append(body, body1, body2)
	successRes := response.ClientResponse(http.StatusOK, "the request was succesful", body, nil)
	ctx.JSON(http.StatusOK, successRes)
}

// ----------------------------------------------------- get wallet by userid ---------------------------------------- \\

// @Summary Wallet of User
// @Description Get Wallet By User Id
// @Tags Wallet Management
// @Produce json
// @Param user_id query int true "user_id"
// @Security BearerTokenAuth
// @Success 200 {object} response.Response "success"
// @Failure 500 {object} response.Response{} "fail"
// @Router /users/wallet/view [get]
func (handler *OrderHandler) ViewWallet(ctx *gin.Context) {
	userId := ctx.Query("user_id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadGateway, "error in reading the user id", nil, err.Error())
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}

	body, err := handler.orderUseCase.ViewWalletByUserId(userIdInt)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadGateway, "error while getting wallet", nil, err.Error())
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "the request was succesful", body, nil)
	ctx.JSON(http.StatusOK, successRes)
}

// @Summary Print Invoice
// @Description Print invoice by order Id
// @Tags Order Management
// @Produce json
// @Param order_id query int true "order_id"
// @Security BearerTokenAuth
// @Success 200 {object} response.Response "success"
// @Failure 500 {object} response.Response{} "fail"
// @Router /users/order/print [get]
func (handler *OrderHandler) PrintInvoice(ctx *gin.Context) {
	orderId := ctx.Query("order_id")
	orderIdInt, err := strconv.Atoi(orderId)
	if err != nil {
		err = errors.New("error in converting order id" + err.Error())
		errRes := response.ClientResponse(http.StatusBadGateway, "error in reading the order id", nil, err)
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}
	pdf, err := handler.orderUseCase.PrintInvoice(orderIdInt)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadGateway, "error in printing invoice", nil, err)
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}
	ctx.Header("Content-Disposition", "attachment;filename=invoice.pdf")

	pdfFilePath := "salesReport/invoice.pdf" //generate temp file path for pdf

	err = pdf.OutputFileAndClose(pdfFilePath)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadGateway, "error in printing invoice", nil, err)
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}

	// set header for the file download
	ctx.Header("Content-Disposition", "attachment; filename=sales_report.pdf")
	ctx.Header("Content-Type", "application/pdf")

	//serve pdf file for download
	ctx.File(pdfFilePath)

	//set content type header to application/pdf
	ctx.Header("Content-Type", "application/pdf")

	//write pdf data to the response writer
	err = pdf.Output(ctx.Writer)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadGateway, "error in printing invoice", nil, err)
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "the request was succesful", pdf, nil)
	// ctx.File("invoice.pdf")
	ctx.JSON(http.StatusOK, successRes)
}
