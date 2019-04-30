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
	})

})
