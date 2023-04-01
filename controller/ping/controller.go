package ping

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"

	"github.com/ahdark-services/ahdark-me-redirector/service/blog"
)

var tracer = otel.Tracer("controller.ping")

type Controller struct {
	fx.In
	BlogSvc blog.Service
}

func RegisterController(ctx context.Context, r *gin.Engine, c Controller) {
	ctx, span := tracer.Start(ctx, "register-controller")
	defer span.End()

	r.GET("/ping", c.PingHandler)
}
