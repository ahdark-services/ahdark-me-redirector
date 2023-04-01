package controller

import (
	"go.uber.org/fx"

	"github.com/ahdark-services/ahdark-me-redirector/controller/blog"
	"github.com/ahdark-services/ahdark-me-redirector/controller/ping"
	"github.com/ahdark-services/ahdark-me-redirector/controller/redirects"
)

func Module() fx.Option {
	return fx.Module("controller",
		fx.Invoke(redirects.RegisterController),
		fx.Invoke(blog.RegisterController),
		fx.Invoke(ping.RegisterController),
	)
}
