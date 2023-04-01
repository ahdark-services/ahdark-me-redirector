package env

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

func NewConfig(ctx context.Context) (*Config, error) {
	ctx, span := tracer.Start(ctx, "new-config")
	defer span.End()

	logger := logrus.WithContext(ctx)

	if err := godotenv.Load(".env"); err != nil {
		logger.WithError(err).Warn("failed to load .env file")
		// ignore error
	}

	var systemConfig system
	if err := envconfig.Process("SYSTEM", &systemConfig); err != nil {
		logger.WithError(err).Error("failed to process system config")
		return nil, err
	}

	var logConfig log
	if err := envconfig.Process("LOG", &logConfig); err != nil {
		logger.WithError(err).Error("failed to process log config")
		return nil, err
	}

	var serverConfig server
	if err := envconfig.Process("SERVER", &serverConfig); err != nil {
		logger.WithError(err).Error("failed to process server config")
		return nil, err
	}

	var redisConfig redis
	if err := envconfig.Process("REDIS", &redisConfig); err != nil {
		logger.WithError(err).Error("failed to process redis config")
		return nil, err
	}

	var corsConfig cors
	if err := envconfig.Process("CORS", &corsConfig); err != nil {
		logger.WithError(err).Error("failed to process cors config")
		return nil, err
	}

	var customConfig custom
	if err := envconfig.Process("CUSTOM", &customConfig); err != nil {
		logger.WithError(err).Error("failed to process custom config")
		return nil, err
	}

	return &Config{
		System: systemConfig,
		Log:    logConfig,
		Server: serverConfig,
		Redis:  redisConfig,
		Cors:   corsConfig,
		Custom: customConfig,
	}, nil
}
