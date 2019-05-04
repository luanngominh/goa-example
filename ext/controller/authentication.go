package controller

import (
	"github.com/goadesign/goa"
	"github.com/luanngominh/goa-example/app"
)

// AuthenticationController implements the authentication resource.
type AuthenticationController struct {
	*goa.Controller
}

// NewAuthenticationController creates a authentication controller.
func NewAuthenticationController(service *goa.Service) *AuthenticationController {
	return &AuthenticationController{Controller: service.NewController("AuthenticationController")}
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
	// AuthenticationController_Register: start_implement

	// Put your logic here

	res := &app.Token{}
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
