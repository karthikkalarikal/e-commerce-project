package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/karthikkalarikal/ecommerce-project/pkg/usecase/interfaces"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
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
// @Param cart_id query int false "cart_id"
// @Param user_id query int true "user_id"
// @Param product body models.CartItems true "Cart details"
// @Security BearerTokenAuth
// @Success 200 {object} response.Response "success"
// @Failure 500 {object} response.Response{} "fail"
// @Router /users/carts/addtocart [post]
func (handler *CartHandler) AddToCart(ctx *gin.Context) {
	var cart models.CartItems
	userId := ctx.Query("user_id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in user id", nil, err.Error())
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}

	cartId := ctx.Query("cart_id")
	cartIdInt := 0
	if cartId != "" {
		cartIdInt, err = strconv.Atoi(cartId)
		if err != nil {
			cartIdInt = 0
		}
	}

	if err := ctx.BindJSON(&cart); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields are in wrong format", nil, err.Error())
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}

	resCart, err := handler.cartUsecase.AddToCart(cart, userIdInt, cartIdInt)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields are in wrong format", nil, err.Error())
		ctx.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "succesfully added to cart", resCart, nil)
	ctx.JSON(http.StatusOK, successRes)

}

// CartItemListing godoc
// @Summary List the products in cart
// @Description Retrive and display product list in cart
// @Tags Cart Mangement
// @Produce json
// @Param user_id query int true "user id"
// @Security BearerTokenAuth
// @Success 200 {array} response.Response "Array of product details "
// @Failure 400 {array} response.Response "Bad request"
// @Router /users/carts/viewcart [get]
func (handler *CartHandler) CartItemListing(c *gin.Context) {
	fmt.Println("*************cart item listing****************")
	user_id := c.Query("user_id")
	userInt, err := strconv.Atoi(user_id)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	cartItems, err := handler.cartUsecase.CartItemListing(userInt)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not display the  products", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "succesfully displayed the products in cart", cartItems, nil)
	c.JSON(http.StatusOK, successRes)
}

// CartItemQuatityModification godoc
// @Summary Update the quantity of cart
// @Description Change the quantity of the cart items
// @Tags Cart Mangement
// @Produce json
// @Param user_id query int true "user id"
// @Param product_id query int true "product_id"
// @Param quantity query string true "quantity"
// @Security BearerTokenAuth
// @Success 200 {array} response.Response "Array of product details "
// @Failure 400 {array} response.Response "Bad request"
// @Router /users/carts/quantity [patch]
func (handler *CartHandler) CartItemQuatityModification(c *gin.Context) {
	fmt.Println("*************cart item mod****************")
	user_id := c.Query("user_id")
	userInt, err := strconv.Atoi(user_id)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in user id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	product_id := c.Query("product_id")
	productInt, err := strconv.Atoi(product_id)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in product id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	quantity := c.Query("quantity")
	cartItems, err := handler.cartUsecase.CartItemQuantityUpdations(userInt, productInt, quantity)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not display the  products", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "succesfully changed the quantity", cartItems, nil)
	c.JSON(http.StatusOK, successRes)
}

// CartItemDeletion godoc
// @Summary delete products from cart
// @Description delet a product from cart
// @Tags Cart Mangement
// @Produce json
// @Param user_id query int true "user id"
// @Param product_id query int true "product_id"
// @Security BearerTokenAuth
// @Success 200 {array} response.Response "Array of product details "
// @Failure 400 {array} response.Response "Bad request"
// @Router /users/carts/delete [delete]
func (handler *CartHandler) CartItemDeletion(c *gin.Context) {
	fmt.Println("*************cart item mod****************")
	user_id := c.Query("user_id")
	userInt, err := strconv.Atoi(user_id)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in user id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	product_id := c.Query("product_id")
	productInt, err := strconv.Atoi(product_id)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in product id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	cartItems, err := handler.cartUsecase.CartItemDeletion(userInt, productInt)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not delete", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "succesfully deleted the product", cartItems, nil)
	c.JSON(http.StatusOK, successRes)
}
