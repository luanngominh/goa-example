//go:generate goagen bootstrap -d github.com/luanngominh/goa-example/goa/design

package main

import (
	"os"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/joho/godotenv"

	"github.com/luanngominh/goa-example/app"
	. "github.com/luanngominh/goa-example/ext/controller"
)

func main() {
	if os.Getenv("ENV") == "DEV" {
		if err := godotenv.Load(); err != nil {
			panic(err)
		}
	}

	// Create service
	service := goa.New("Take Note Backend API")

	// Setup jwt middleware
	jwtMiddleware, err := jwt.NewJWTMiddleware(app.NewJWTSecurity)
	if err != nil {
		panic(err)
	}

	// Mount middleware
	// Mount jwt middleware
	app.UseJWTMiddleware(service, jwtMiddleware)

	// Mount default middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Create db Service
	dbSvc, dbClose := db.New(os.Getenv("DB_CONNECTION"))

	// Mount "authentication" controller
	c := NewAuthenticationController(service, dbSvc)
	app.MountAuthenticationController(service, c)

	c1 := NewHealthCheckController(service, dbSvc)
	app.MountHealthCheckController(service, c1)

	// Start service
	if err := service.ListenAndServe(":8888"); err != nil {
		service.LogError("startup", "err", err)
	}
}
