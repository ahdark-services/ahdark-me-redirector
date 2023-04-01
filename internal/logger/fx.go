package logger

import (
	"github.com/sirupsen/logrus"
	fxlogrus "github.com/takt-corp/fx-logrus"
	"go.uber.org/fx/fxevent"
)

func FxLogger(logger *logrus.Logger) fxevent.Logger {
	return &fxlogrus.LogrusLogger{
		Logger: logger,
	}
}
