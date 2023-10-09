package usecase

import (
	"errors"
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/karthikkalarikal/ecommerce-project/pkg/config"
	"github.com/karthikkalarikal/ecommerce-project/pkg/helper"
	"github.com/karthikkalarikal/ecommerce-project/pkg/repository/interfaces"
	usecase "github.com/karthikkalarikal/ecommerce-project/pkg/usecase/interfaces"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
	"github.com/twilio/twilio-go"
	verify "github.com/twilio/twilio-go/rest/verify/v2"
)

type otpUseCaseImpl struct {
	cfg      config.Config
	otpRepo  interfaces.OtpRepository
	userRepo interfaces.UserRepository
}

func NewOtpUsecase(repo interfaces.OtpRepository, cfg config.Config, repoUser interfaces.UserRepository) usecase.OtpUseCase {
	return &otpUseCaseImpl{
		otpRepo:  repo,
		cfg:      cfg,
		userRepo: repoUser,
	}
}

// -----------------------------verify otp------------------------------------\\
func (otp *otpUseCaseImpl) VerifyOTP(code models.VerifyData) (models.TokenUsers, error) {
	twilioClient := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: otp.cfg.ACCOUNTSID,
		Password: otp.cfg.AUTHTOKEN,
	})
	params := &verify.CreateVerificationCheckParams{}
	params.SetTo("+91" + code.PhoneNumber)
	params.SetCode(code.Code)
	resp, err := twilioClient.VerifyV2.CreateVerificationCheck(otp.cfg.SERVICESID, params)

	if err != nil {
		return models.TokenUsers{}, errors.New("could not verify the phone number")
	}

	if *resp.Status != "approved" {
		return models.TokenUsers{}, errors.New("could not verifu the token user")
	}

	userDetails, err := otp.otpRepo.UserDetailsUsingPhone(code.PhoneNumber)
	if err != nil {
		return models.TokenUsers{}, err
	}

	var user models.UserDetailsResponse
	err = copier.Copy(&user, &userDetails)
	if err != nil {
		return models.TokenUsers{}, err
	}

	tokenString, err := helper.GenerateTokenClients(userDetails)
	if err != nil {
		return models.TokenUsers{}, err
	}

	return models.TokenUsers{
		Users: user,
		Token: tokenString,
	}, nil
}

// -------------------------------------------send otp ------------------------------------------\\
func (otp *otpUseCaseImpl) SendOTP(phone string) error {
	fmt.Println("i am here")
	user, err := otp.otpRepo.FindUserByMobileNumber(phone)
	if err != nil {
		return errors.New("the user does not exist")
	}
	ok, err := otp.userRepo.UserBlockedStatus(user.Email)
	if err != nil {
		return errors.New("error gettin user status")
	}
	if ok {
		return errors.New("the user is blocked by admin")
	}

	twilioClient := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: otp.cfg.ACCOUNTSID,
		Password: otp.cfg.AUTHTOKEN,
	})

	to := "+91" + phone
	params := &verify.CreateVerificationParams{}
	params.SetTo(to)
	params.SetChannel("sms")
	fmt.Println("service id", otp.cfg.SERVICESID, otp.cfg.AUTHTOKEN, otp.cfg.ACCOUNTSID, otp.cfg)
	fmt.Println("params", params)
	_, err = twilioClient.VerifyV2.CreateVerification(otp.cfg.SERVICESID, params)
	if err != nil {
		fmt.Println(err)

		return errors.New("error occured while generation otp")

	}
	return nil
}
