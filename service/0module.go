package service

import (
	"go.uber.org/fx"

	"github.com/ahdark-services/ahdark-me-redirector/service/blog"
)

func Module() fx.Option {
	return fx.Module("service",
		fx.Provide(blog.NewService),
	)
}
