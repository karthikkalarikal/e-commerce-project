package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
	"github.com/karthikkalarikal/ecommerce-project/pkg/usecase/interfaces"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
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

// AddProduct is a function to add a new product by admin.
// @Summary Add product
// @Description Add product by admin
// @Tags Product Management
// @Accept json
// @Produce json
// @Param product body models.Product true "Product object"
// @Security BearerTokenAuth
// @Success 200 {object}  response.Response "Added product details"
// @Failure 400 {object}  response.Response "Bad request"
// @Router /admin/product/addproduct [post]
func (u *ProductHandler) AddProduct(c *gin.Context) {
	var product models.Product

	if err := c.BindJSON(&product); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	returnProduct, err := u.productUsecase.AddProduct(product)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Could not add the product", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully added the product", returnProduct, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary Delete product
// @Description Delete product by id
// @Tags Product Management
// @Accept json
// @Produce json
// @Param product_id query int true "product id"
// @Security BearerTokenAuth
// @Success 200 {object} response.Response "Array of deleted product details "
// @Failure 400 {object} response.Response  "Bad request"
// @Router /admin/product/deleteproduct/{product_id} [delete]
func (u *ProductHandler) DeleteProduct(c *gin.Context) {

	id_str := c.Query("product_id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "id is in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	product, err := u.productUsecase.DeleteProduct(id)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in deleting the product", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	succesRes := response.ClientResponse(http.StatusOK, "sucessfully deleted the product", product, nil)
	c.JSON(http.StatusOK, succesRes)
}

// GetProductList godoc
// @Summary List the users you could specify page and no of products in one page
// @Description Retrive and display product list according to instructions
// @Tags General
// @Produce json
// @Param page query int false "Page number (default 1)"
// @Param per_page query int false "Results per page (default 10)"
// @Success 200 {array} response.Response "Array of product details "
// @Failure 400 {array} response.Response "Bad request"
// @Router /users/viewproducts [get]
func (u *ProductHandler) ListProducts(c *gin.Context) {
	fmt.Println("list product handler")

	pageNo := c.DefaultQuery("page", "1")       // default 1
	pageList := c.DefaultQuery("per_page", "5") // default to 5
	pageNoInt, err := strconv.Atoi(pageNo)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Products cannot be displayed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	pageListInt, err := strconv.Atoi(pageList)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Products cannot be displayed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	product_list, err := u.productUsecase.ListProducts(pageNoInt, pageListInt)
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
// @Security BearerTokenAuth
// @Success 200 {array} domain.Category "Update Category  "
// @Failure 400 {array} domain.Category  "Bad request"
// @Router /admin/product/updatecategory [post]
func (u *ProductHandler) UpdateCategory(c *gin.Context) {
	var category domain.Category

	id_str := c.Param("id")
	id, err := strconv.Atoi(id_str)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	if err := c.BindJSON(&category); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	body, err := u.productUsecase.UpdateCategory(category, id)
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
// @Param category_id query int true "category_id"
// @Security BearerTokenAuth
// @Success 200 {array} response.Response "delete Category  "
// @Failure 400 {array} response.Response  "Bad request"
// @Router /admin/product/deletecategory/{category_id} [delete]
func (u *ProductHandler) DeleteCategory(c *gin.Context) {

	id_str := c.Query("category_id")
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

	succesRes := response.ClientResponse(http.StatusOK, "deleted a category successfully", category, nil)
	c.JSON(http.StatusOK, succesRes)

}

// AddCategory is a function to add a new category by admin.
// @Summary Add category
// @Description Add category by admin
// @Tags Product Management
// @Accept json
// @Produce json
// @Param product body models.Category true "Category object"
// @Security BearerTokenAuth
// @Success 200 {object}  response.Response "Added Category details"
// @Failure 400 {object}  response.Response"Bad request"
// @Router /admin/product/addcategory [post]
func (u *ProductHandler) AddCategory(c *gin.Context) {
	var adCat domain.Category

	if err := c.BindJSON(&adCat); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	returnCategory, err := u.productUsecase.AddCategory(adCat)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Could not add the category", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully added the category", returnCategory, nil)
	c.JSON(http.StatusOK, successRes)
}

// EditProduct is a function to edit product by admin.
// @Summary Edit product
// @Description Edit product by admin
// @Tags Product Management
// @Accept json
// @Produce json
// @Param product_id query int true "product_id"
// @Param product body models.Product true "Product object"
// @Security BearerTokenAuth
// @Success 200 {string}  response.Response "Edit product details"
// @Failure 400 {string}  response.Response "Bad request"
// @Router /admin/product/editproduct/{product_id} [put]
func (u *ProductHandler) EditProduct(c *gin.Context) {
	var product domain.Product

	id := c.Query("product_id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "problems in the id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	if err := c.BindJSON(&product); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields are in the wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	modProduct, err := u.productUsecase.EditProduct(product, idInt)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not edit the product", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	succesRes := response.ClientResponse(http.StatusOK, "sucessfully edited product", modProduct, nil)
	c.JSON(http.StatusOK, succesRes)
}
