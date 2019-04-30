//go:generate goagen bootstrap -d github.com/luanngominh/goa-example/goa/design

package main

import (
	"os"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"

	"github.com/luanngominh/goa-example/app"
	. "github.com/luanngominh/goa-example/ext/controller"
	dbsvc "github.com/luanngominh/goa-example/ext/service"
)

func main() {
	if os.Getenv("ENV") == "DEV" {
		// load .env file
	}
	
	// Create service
	service := goa.New("Take Note Backend API")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Create db Service
	dbSvc := &dbsvc.Service{}

	// Mount "authentication" controller
	c := NewAuthenticationController(service, dbSvc)
	app.MountAuthenticationController(service, c)

	// Start service
	if err := service.ListenAndServe(":8888"); err != nil {
		service.LogError("startup", "err", err)
	}
}
