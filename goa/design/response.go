package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

//LoginResponse response for login endpoint
var LoginResponse = MediaType("application/vnd.token+json", func() {
	Description("Login reponse")

	Attribute("token", String)

	View("default", func() {
		Attribute("token")
	})
})
