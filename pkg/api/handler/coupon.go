package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	usecase "github.com/karthikkalarikal/ecommerce-project/pkg/usecase/interfaces"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/response"
)

type CouponHandler struct {
	usecase usecase.CouponUseCase
}

func NewCouponHandler(usecase usecase.CouponUseCase) *CouponHandler {
	return &CouponHandler{
		usecase: usecase,
	}
}

// ----------------------------------------- add coupon ----------------------------------------------------\\

// @Summary Add Coupon
// @Description Add Coupon
// @Tags Coupon Mangement
// @Accept json
// @Produce json
// @Param product body models.CouponInput true "coupon details"
// @Security BearerTokenAuth
// @Success 200 {object} response.Response "success"
// @Failure 500 {object} response.Response{} "fail"
// @Router /admin/coupon/addcoupon [post]
func (hander *CouponHandler) AddCoupon(c *gin.Context) {
	var body models.CouponInput

	if err := c.BindJSON(&body); err != nil {
		err = errors.New("error whild binding the json to struct" + err.Error())
		errRes := response.ClientResponse(http.StatusBadRequest, "fields are in the wrong format", nil, err)
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	body2, err := hander.usecase.AddCoupon(body.Coupon, body.DiscountPercentage, body.Validity, body.MinimumPrice)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not add the coupon", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "succesfully added to coupon", body2, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Vidw Coupons
// @Description View Coupons by Admin
// @Tags Coupon Mangement
// @Produce json
// @Security BearerTokenAuth
// @Success 200 {object} response.Response "success"
// @Failure 500 {object} response.Response{} "fail"
// @Router /admin/coupon/view [get]
func (hander *CouponHandler) ViewCoupon(c *gin.Context) {

	body, err := hander.usecase.ViewCoupon()
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not view coupon", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "succesfully accessed the coupon", body, nil)
	c.JSON(http.StatusOK, successRes)

}
