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
// @Param id path string true "product-id"
// @Security BearerTokenAuth
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response{}
// @Router /cart/addtocart/{id} [put]
func (handler *CartHandler) AddToCart(ctx *gin.Context) {
	var cart domain.Cart
	id_str := ctx.Param("id")
	id, err := strconv.Atoi(id_str)
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

	resCart, err := handler.cartUsecase.AddToCart(cart, id)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields are in wrong format", nil, err.Error())
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "succesfully added to cart", resCart, nil)
	ctx.JSON(http.StatusOK, successRes)

}

// Cart item listing
