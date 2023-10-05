package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/karthikkalarikal/ecommerce-project/pkg/usecase/interfaces"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/response"
)

type ProductHandler struct {
	productUsecase interfaces.ProductUseCase
}

func NewProductHandler(usecase interfaces.ProductUseCase) *ProductHandler {
	return &ProductHandler{
		productUsecase: usecase,
	}
}

// @Summary ViewProducts
// @Description view products by a user
// @Accept json
// @Produce json
// @Success 200 {object} models.Product "List of products"
// @Failure 400 {array} models.Product "Bad request"
// @Router /user/viewproducts [post]
func (u *ProductHandler) ListProducts(c *gin.Context) {
	fmt.Println("list product handler")

	product_list, err := u.productUsecase.ListProducts()
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Products cannot be displayed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	message := "products list"

	successRes := response.ClientResponse(http.StatusOK, message, product_list, nil)
	// fmt.Println(product_list)
	c.JSON(http.StatusOK, successRes)
}
