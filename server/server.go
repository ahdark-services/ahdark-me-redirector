package server

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"

	"github.com/ahdark-services/ahdark-me-redirector/internal/env"
	"github.com/ahdark-services/ahdark-me-redirector/middleware"
)

func NewServer(ctx context.Context, tp *tracesdk.TracerProvider) (*gin.Engine, error) {
	ctx, span := tracer.Start(ctx, "new-server")
	defer span.End()

	r := gin.New()
	r.Use(otelgin.Middleware("server.gin", otelgin.WithTracerProvider(tp)))

	return r, nil
}

func ConfigServer(ctx context.Context, r *gin.Engine, config *env.Config) {
	ctx, span := tracer.Start(ctx, "config-server")
	defer span.End()

	if config.System.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r.Use(middleware.CORS(config))
	r.Use(middleware.Gzip())
	r.Use(middleware.RequestID())
}
