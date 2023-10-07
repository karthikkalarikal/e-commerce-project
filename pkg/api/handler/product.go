package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
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
// @Router /users/viewproducts [get]
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

// @Summary update category
// @Description Category by id
// @Tags Product Management
// @Accept json
// @Produce json
//
//	@Param id query int true "category_id"
//
// @Security ApiKeyHeaderAuth
// @Success 200 {array} domain.Category "Update Category  "
// @Failure 400 {array} domain.Category  "Bad request"
// @Router /admin/product/updatecategory [patch]
func (u *ProductHandler) UpdateCategory(c *gin.Context) {
	var category domain.Category

	if err := c.BindJSON(&category); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	body, err := u.productUsecase.UpdateCategory(category)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Could not update the category", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully updated the category", body, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary delete category
// @Description delete Category by id
// @Tags Product Management
// @Accept json
// @Produce json
//
//	@Param id query int true "category_id"
//
// @Security ApiKeyHeaderAuth
// @Success 200 {array} domain.Category "delete Category  "
// @Failure 400 {array} domain.Category  "Bad request"
// @Router /admin/product/updatecategory [patch]
func (u *ProductHandler) DeleteCategory(c *gin.Context) {

	id_str := c.Query("id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in converting the id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	category, err := u.productUsecase.DeleteCategory(id)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could delete category", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	succesRes := response.ClientResponse(http.StatusOK, "added category successfully", category, nil)
	c.JSON(http.StatusOK, succesRes)

}
