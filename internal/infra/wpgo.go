package infra

import (
	"context"

	"github.com/ahdark-services/ahdark-me-redirector/internal/env"
	"github.com/ahdark-services/ahdark-me-redirector/pkg/wpgo"
)

func NewWpGoClient(ctx context.Context, config *env.Config) *wpgo.Client {
	ctx, span := tracer.Start(ctx, "new-wpgo-client")
	defer span.End()

	if config.Custom.WordPressEndpoint == "" {
		return nil
	}

	return wpgo.NewClient(config.Custom.WordPressEndpoint)
}
