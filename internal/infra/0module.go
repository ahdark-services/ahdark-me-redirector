package infra

import (
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"
)

var tracer = otel.Tracer("internal.infra")

func Module() fx.Option {
	return fx.Module("internal.infra",
		fx.Provide(NewRedisClient),
		fx.Provide(NewCacheDriver),
	)
}
