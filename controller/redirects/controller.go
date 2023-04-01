package redirects

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"

	"github.com/ahdark-services/ahdark-me-redirector/internal/env"
	"github.com/ahdark-services/ahdark-me-redirector/pkg/util"
)

var tracer = otel.Tracer("controller.redirects")

type Controller struct {
	fx.In
	Config *env.Config
}

func RegisterController(ctx context.Context, r *gin.Engine, c Controller) {
	ctx, span := tracer.Start(ctx, "register-controller")
	defer span.End()

	if c.Config.Custom.RedirectConfig == "" {
		logrus.WithContext(ctx).Warn("no redirects configured, skipping registration")
		return
	}

	file, err := os.ReadFile(util.AbsolutePath(c.Config.Custom.RedirectConfig))
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Fatal("failed to read redirect config")
		return
	}

	var config map[string]string
	if err := json.Unmarshal(file, &config); err != nil {
		logrus.WithContext(ctx).WithError(err).Fatal("failed to unmarshal redirect config")
		return
	}

	for key, u := range config {
		path := fmt.Sprintf("/%s", key)
		r.Any(path, c.RedirectHandler(u))
		logrus.WithContext(ctx).WithField("path", path).WithField("url", u).Info("registered redirect")
	}
}
