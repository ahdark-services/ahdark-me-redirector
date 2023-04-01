package entry

import (
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"

	"github.com/ahdark-services/ahdark-me-redirector/controller"
	"github.com/ahdark-services/ahdark-me-redirector/internal/env"
	"github.com/ahdark-services/ahdark-me-redirector/internal/infra"
	"github.com/ahdark-services/ahdark-me-redirector/internal/logger"
	"github.com/ahdark-services/ahdark-me-redirector/internal/trace"
	"github.com/ahdark-services/ahdark-me-redirector/server"
)

var tracer = otel.Tracer("entry")

func Entry() []fx.Option {
	return []fx.Option{
		fx.Provide(env.NewConfig),
		fx.Provide(logger.NewLogger),
		fx.WithLogger(logger.FxLogger),
		trace.Module(),
		infra.Module(),
		server.Module(),
		controller.Module(),
	}
}
