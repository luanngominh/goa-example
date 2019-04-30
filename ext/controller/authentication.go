package controller

import (
	"github.com/goadesign/goa"

	"github.com/luanngominh/goa-example/app"
	"github.com/luanngominh/goa-example/ext/services"
)

// AuthenticationController implements the authentication resource.
type AuthenticationController struct {
	*goa.Controller
	Service *services.Services
}

// NewAuthenticationController creates a authentication controller.
func NewAuthenticationController(service *goa.Service, dbSvc *services.Services) *AuthenticationController {
	return &AuthenticationController{
		Controller: service.NewController("AuthenticationController"),
		Service: dbSvc,
	}
}

// Login runs the login action.
func (c *AuthenticationController) Login(ctx *app.LoginAuthenticationContext) error {
	// AuthenticationController_Login: start_implement

	// Put your logic here

	res := &app.Token{}
	return ctx.OK(res)
	// AuthenticationController_Login: end_implement
}
