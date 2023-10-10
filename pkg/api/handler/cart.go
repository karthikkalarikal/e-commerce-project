package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
	"github.com/karthikkalarikal/ecommerce-project/pkg/usecase/interfaces"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/response"
)

type CartHandler struct {
	cartUsecase interfaces.CartUseCase
}

func NewCartHandler(usecase interfaces.CartUseCase) *CartHandler {
	return &CartHandler{
		cartUsecase: usecase,
	}
}

// @Summary Add to Cart
// @Description Add product to the cart using product id
// @Tags Cart Mangement
// @Accept json
// @Produce json
// @Param user_id query int true "user_id"
// @Param product_id query int true "product_id"
// @Param product body models.Cart true "Cart details"
// @Security BearerTokenAuth
// @Success 200 {object} response.Response "success"
// @Failure 500 {object} response.Response{} "fail"
// @Router /users/user/addtocart [post]
func (handler *CartHandler) AddToCart(ctx *gin.Context) {
	var cart domain.Cart
	userId := ctx.Query("user_id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in user id", nil, err.Error())
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}

	productId := ctx.Query("product_id")
	productIdInt, err := strconv.Atoi(productId)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in user id", nil, err.Error())
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}

	if err := ctx.BindJSON(&cart); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields are in wrong format", nil, err.Error())
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}

	resCart, err := handler.cartUsecase.AddToCart(cart, userIdInt, productIdInt)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields are in wrong format", nil, err.Error())
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "succesfully added to cart", resCart, nil)
	ctx.JSON(http.StatusOK, successRes)

}

// cart item listing
