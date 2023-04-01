package infra

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"

	"github.com/ahdark-services/ahdark-me-redirector/internal/env"
)

func NewRedisClient(ctx context.Context, config *env.Config) *redis.Client {
	ctx, span := tracer.Start(ctx, "new-redis-client")
	defer span.End()

	logger := logrus.WithContext(ctx)

	if config.System.CacheDriver != "redis" {
		logger.Warn("redis is not configured as cache driver")
		return nil
	}

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
		Username: config.Redis.Username,
		Password: config.Redis.Password,
		DB:       config.Redis.DB,
	})

	if err := redisotel.InstrumentTracing(client); err != nil {
		logger.WithError(err).Warn("failed to instrument tracing")
		// ignore error
	}

	return client
}
