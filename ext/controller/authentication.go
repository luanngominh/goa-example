package controller

import (
	"github.com/goadesign/goa"

	"github.com/luanngominh/goa-example/app"
	"github.com/luanngominh/goa-example/ext/service"
	"github.com/luanngominh/goa-example/ext/util/crypto"
	"github.com/luanngominh/goa-example/model"
)

// AuthenticationController implements the authentication resource.
type AuthenticationController struct {
	*goa.Controller
	Service *service.Service
}

// NewAuthenticationController creates a authentication controller.
func NewAuthenticationController(service *goa.Service, dbSvc *service.Service) *AuthenticationController {
	return &AuthenticationController{
		Controller: service.NewController("AuthenticationController"),
		Service:    dbSvc,
	}
}

// ForgotPassword runs the forgot_password action.
func (c *AuthenticationController) ForgotPassword(ctx *app.ForgotPasswordAuthenticationContext) error {
	// AuthenticationController_ForgotPassword: start_implement

	// Put your logic here

	res := &app.ForgotPassword{}
	return ctx.OK(res)
	// AuthenticationController_ForgotPassword: end_implement
}

// Login runs the login action.
func (c *AuthenticationController) Login(ctx *app.LoginAuthenticationContext) error {
	// AuthenticationController_Login: start_implement

	// Put your logic here

	res := &app.Token{}
	return ctx.OK(res)
	// AuthenticationController_Login: end_implement
}

// Refresh runs the refresh action.
func (c *AuthenticationController) Refresh(ctx *app.RefreshAuthenticationContext) error {
	// AuthenticationController_Refresh: start_implement

	// Put your logic here

	res := &app.Refresh{}
	return ctx.OK(res)
	// AuthenticationController_Refresh: end_implement
}

// Register runs the register action.
func (c *AuthenticationController) Register(ctx *app.RegisterAuthenticationContext) error {
	payload := ctx.Payload

	passwordDigest, err := crypto.BcryptHashPassword(payload.Password)
	if err != nil {
		return err
	}

	u, err := c.Service.User.Create(ctx.Context, &model.User{
		Email: payload.Email,
		Fullname: payload.Fullname,
		PasswordDigest: passwordDigest,
	})

	

	if err != nil {
		return err
	}

	res := &app.Register{

	}
	return ctx.OK(res)
	// AuthenticationController_Register: end_implement
}

// UpdateInformation runs the update_information action.
func (c *AuthenticationController) UpdateInformation(ctx *app.UpdateInformationAuthenticationContext) error {
	// AuthenticationController_UpdateInformation: start_implement

	// Put your logic here

	res := &app.UpdateUser{}
	return ctx.OK(res)
	// AuthenticationController_UpdateInformation: end_implement
}

// VerifyCode runs the verify_code action.
func (c *AuthenticationController) VerifyCode(ctx *app.VerifyCodeAuthenticationContext) error {
	// AuthenticationController_VerifyCode: start_implement

	// Put your logic here

	res := &app.Verify{}
	return ctx.OK(res)
	// AuthenticationController_VerifyCode: end_implement
}
