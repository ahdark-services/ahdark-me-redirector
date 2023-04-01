package main

import (
	"context"

	"go.opentelemetry.io/otel"

	"github.com/ahdark-services/ahdark-me-redirector/entry"
)

var (
	tracer = otel.Tracer("main")
	ctx    = context.Background()
)

func main() {
	ctx, span := tracer.Start(ctx, "main")
	defer span.End()

	entry.Bootstrap(ctx)
}
