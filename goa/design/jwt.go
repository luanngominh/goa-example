package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

//JWT define api:access scope
var JWT = JWTSecurity("jwt", func() {
	Description("Use JWT to authentication")
	Header("Authorization")
	Scope("api:access", "API access")
})