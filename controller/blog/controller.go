package blog

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"

	"github.com/ahdark-services/ahdark-me-redirector/internal/env"
	"github.com/ahdark-services/ahdark-me-redirector/pkg/wpgo"
)

var tracer = otel.Tracer("controller.redirects")

type Controller struct {
	fx.In
	Config *env.Config
	WpGo   *wpgo.Client `optional:"true"`
}

func RegisterController(ctx context.Context, r *gin.Engine, c Controller) {
	ctx, span := tracer.Start(ctx, "register-controller")
	defer span.End()

	if c.Config.Custom.WordPressEndpoint == "" {
		logrus.WithContext(ctx).Info("wordpress endpoint is not set, skipping blog controller")
		return
	}

	r.GET("/blog/:id", c.RedirectPostHandler)
	r.HEAD("/blog/:id", c.RedirectPostHandler)
	r.OPTIONS("/blog/:id", c.RedirectPostHandler)
}
