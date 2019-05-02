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

// var RefreshResponse = MediaType("application/vnd.refresh+json", func() {
// 	Description("Refresh token response")

// 	Attribute("token", String)

// 	View("default", func() {
// 		Attribute("token", String)
// 	})
// })

// var ForgotPasswordResponse = MediaType("application/forgot_password+json", func() {
// 	Description("Forgot password response")

// 	Attribute("status", String)

// 	View("default", func() {
// 		Attribute("status", String)
// 	})
// })

// var VerifyCodeResponse = MediaType("application/vnd.verify+json", func() {
// 	Description("Verify code response")

// 	Attribute("token", String)
// 	Attribute("message", String)

// 	View("default", func() {
// 		Attribute("token")
// 		Attribute("message")
// 	})
// })

// var UpdateUserInformationResponse = MediaType("application/vnd.update_user+json", func() {
// 	Description("Update user information")

// 	Attribute("email", Email)
// 	Attribute("fullname", String)	

// 	View("default", func() {
// 		Attribute("email", Email)
// 		Attribute("fullname", String)
// 	})
// })