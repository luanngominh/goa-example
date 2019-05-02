package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("authentication", func() {
	BasePath("/auth")

	Action("login", func() {
		NoSecurity()
		Routing(
			POST("/login"),
		)
		Description("Sign in")
		Payload(LoginPayload)
		Response(OK, LoginResponse)
		Response(BadRequest)
	})

	Action("register", func() {
		NoSecurity()
		Routing(
			POST("/register"),
		)
		Payload(RegisterPayload)
		Response(OK, LoginResponse)
		Response(BadRequest)
	})
	
	// Action("refresh", func() {
	// 	Security(JWT, func() {
	// 		Scope("api:access")
	// 	})
	// 	Routing(
	// 		POST("/refresh"),
	// 	)
	// 	Payload(RefreshPayLoad)
	// 	Response(OK, RefreshResponse)
	// 	Response(BadRequest)
	// })

	// Action("forgot_password", func() {
	// 	NoSecurity()
	// 	Routing(
	// 		POST("/forgotpass"),
	// 	)
	// 	Payload(ForgotPasswordPayload)
	// 	Response(OK, ForgotPasswordResponse)
	// 	Response(BadRequest)
	// })

	// Action("verify_code", func() {
	// 	NoSecurity()
	// 	Routing(
	// 		POST("/lvc"),
	// 	)
	// 	Payload(VerifyCodePayload)
	// 	Response(OK, VerifyCodeResponse)
	// 	Response(BadRequest)
	// })

	// Action("update_information", func() {
	// 	Security(JWT, func() {
	// 		Scope("api:access")
	// 	})
	// 	Payload(UpdateUserInformationPayload)
	// 	Response(OK, UpdateUserInformationResponse)
	// 	Response(BadRequest)
	// })
})