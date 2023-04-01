package entry

import (
	"context"

	"go.uber.org/fx"
)

func Bootstrap(ctx context.Context) {
	ctx, span := tracer.Start(ctx, "bootstrap")
	defer span.End()

	opts := []fx.Option{
		fx.Supply(fx.Annotate(ctx, fx.As(new(context.Context)))),
	}
	opts = append(opts, Entry()...)

	app := fx.New(opts...)
	app.Run()
}
