package trace

import (
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"
)

var tracer = otel.Tracer("internal.trace")

func Module() fx.Option {
	return fx.Module("trace",
		fx.Provide(NewTracerProvider),
		fx.Invoke(UseTracerProvider),
	)
}
