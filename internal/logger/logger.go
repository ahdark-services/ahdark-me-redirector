package logger

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"

	"github.com/ahdark-services/ahdark-me-redirector/internal/env"
	"github.com/ahdark-services/ahdark-me-redirector/pkg/util"
)

var tracer = otel.Tracer("internal.logger")

func NewLogger(ctx context.Context, config *env.Config) (*logrus.Logger, error) {
	ctx, span := tracer.Start(ctx, "new-logger")
	defer span.End()

	logger := logrus.StandardLogger()

	if config.System.Debug {
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}

	switch config.Log.Driver {
	case "console":
		logger.SetOutput(os.Stdout)
		logger.SetFormatter(&logrus.TextFormatter{})
	case "file":
		file, err := os.OpenFile(util.AbsolutePath(config.Log.FilePath), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			logger.WithError(err).Fatal("failed to open log file")
			return nil, err
		}

		logger.SetOutput(file)
		logger.SetFormatter(&logrus.JSONFormatter{})
	default:
		logger.WithField("driver", config.Log.Driver).Fatal("unknown log driver")
	}

	return logger, nil
}
