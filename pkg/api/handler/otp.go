package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	services "github.com/karthikkalarikal/ecommerce-project/pkg/usecase/interfaces"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/response"
)

type OtpHandler struct {
	otpUsecase services.OtpUseCase
}

func NewOtpHandler(useCase services.OtpUseCase) *OtpHandler {
	return &OtpHandler{
		otpUsecase: useCase,
	}
}

// @Summary SendOTP
// @Description verify Phone number using OTP
// @Tags User Auth
// @Accept json
// @Produce json
// @Param user body models.OTPData true "Phone number"
// @Success 200 {array} models.UserSignInResponse "phone number"
// @Failure 400 {array} models.UserSignInResponse "Bad request"
// @Router /users/otplogin [post]
func (otp *OtpHandler) SendOTP(c *gin.Context) {

	var phone models.OTPData
	if err := c.BindJSON(&phone); err != nil {
		errorsRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorsRes)
	}

	err := otp.otpUsecase.SendOTP(phone.PhoneNumber)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not send OTP", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "OTP sent successfully", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary VerifyOTP
// @Description verify Phone number using OTP
// @Tags User Auth
// @Accept json
// @Produce json
// @Param user body models.VerifyData true "Phone number and code"
// @Success 200 {array} models.UserSignInResponse "phone number"
// @Failure 400 {array} models.UserSignInResponse "Bad request"
// @Router /users/verifyotp [post]
func (otp *OtpHandler) VerifyOTP(c *gin.Context) {

	var code models.VerifyData
	if err := c.BindJSON(&code); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	users, err := otp.otpUsecase.VerifyOTP(code)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not verify OTP", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully verified OTP", users, nil)
	c.JSON(http.StatusOK, successRes)

}
