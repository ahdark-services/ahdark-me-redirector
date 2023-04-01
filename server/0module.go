package server

import (
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"
)

var tracer = otel.Tracer("server")

func Module() fx.Option {
	return fx.Module("server",
		fx.Provide(NewServer),
		fx.Invoke(ConfigServer),
		fx.Invoke(HttpListener),
		fx.Invoke(HttpsListener),
		fx.Invoke(UnixListener),
	)
}
