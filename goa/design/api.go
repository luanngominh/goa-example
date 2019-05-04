package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("Take Note Backend API", func() {
	Title("Take Note Backend API")
	Description("An simple API for Take Note")
	Contact(func() {
		Name("Luan Ngo Minh")
		Email("ngominhluanbox@gmail.com")
	})
	Host("localhost")
	Scheme("http")
	BasePath("/api")
	Origin("*", func() {
		Methods("GET", "POST", "PUT", "DELETE", "HEAD")
		Headers("Accept", "Content-Type", "Authorization")
		Expose("Content-Type", "Origin")
		MaxAge(600)
		Credentials()
	})

	Consumes("application/json")
	Produces("application/json")
})
