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

//RegisterPayload ...
var RegisterPayload = Type("ReigsterPayload", func() {
	Attribute("email", String, func() {
		MinLength(5)
		MaxLength(65)
		Format("email")
		Example("me@luanngominh.me")
	})

	Attribute("password", String, func() {
		MinLength(6)
		MaxLength(50)
		Example("secret")
	})

	Attribute("fullname", String, func() {
		MinLength(10)
		MaxLength(100)
		Example("Luan Ngo Minh")
	})

	Required("email", "fullname", "password")
})

//RefreshPayload define refresh payload for refresh token endpoint
var RefreshPayLoad = Type("RefreshPayload", func() {
	Attribute("token", String, func() {
		MaxLength(200)
		Example("xx.xxxx.xx")
	})

	Required("token")
})

var ForgotPasswordPayload = Type("ForgotPassword", func() {
	Attribute("email", String, func() {
		MinLength(5)
		MaxLength(65)
		Format("email")
		Example("me@luanngominh.me")
	})

	Required("email")
})

var VerifyCodePayload = Type("VerifyCodePayload", func() {
	Attribute("email", String, func() {
		MinLength(5)
		MaxLength(65)
		Format("email")
		Example("me@luanngominh.me")
	})

	Attribute("code", String, func() {
		MinLength(6)
		MaxLength(6)
		Example("678512")
	})

	Required("email", "code")
})

var UpdateUserInformationPayload = Type("UpdateUserInformation", func() {
	Attribute("email", String, func() {
		MinLength(5)
		MaxLength(65)
		Format("email")
		Example("me@luanngominh.me")
	})

	Attribute("password", String, func() {
		MinLength(6)
		MaxLength(50)
		Example("secret")
	})

	Attribute("fullname", String, func() {
		MinLength(10)
		MaxLength(100)
		Example("Luan Ngo Minh")
	})
})