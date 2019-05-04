package controller

import (
	"github.com/goadesign/goa"
	"github.com/luanngominh/goa-example/app"
	"github.com/luanngominh/goa-example/ext/service"
)

// HealthCheckController implements the health_check resource.
type HealthCheckController struct {
	*goa.Controller
	Service *service.Service
}

// NewHealthCheckController creates a health_check controller.
func NewHealthCheckController(service *goa.Service, dbSvc *service.Service) *HealthCheckController {
	return &HealthCheckController{
		Controller: service.NewController("HealthCheckController"),
		Service:    dbSvc,
	}
}

// Warm runs the warm action.
func (c *HealthCheckController) Warm(ctx *app.WarmHealthCheckContext) error {
	return ctx.OK([]byte("{}"))
}
