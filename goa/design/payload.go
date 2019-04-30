package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// LoginPayload payload for user login
var LoginPayload = Type("LoginPayload", func() {
	Attribute("email", String, func() {
		MinLength(5)
		MaxLength(65)
		Format("email")
		Example("me@luanngominh.me")
	})

	Attribute("password", String, func() {
		MinLength(6)
		MaxLength(50)
		Example("example-password")
	})

	Required("email", "password")
})
