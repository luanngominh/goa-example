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

var RegisterResponse = MediaType("application/vnd.register+json", func() {
	Description("Register response")

	Attribute("token", String)
	Attribute("message", String)

	View("default", func() {
		Attribute("token")
		Attribute("message")
	})
})