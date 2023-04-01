package infra

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"

	"github.com/ahdark-services/ahdark-me-redirector/internal/env"
	"github.com/ahdark-services/ahdark-me-redirector/pkg/cache"
)

type DriverParams struct {
	fx.In
	RedisClient *redis.Client `optional:"true"`
}

// NewCacheDriver returns a new cache driver
func NewCacheDriver(ctx context.Context, config *env.Config, params DriverParams) (cache.Driver, error) {
	ctx, span := tracer.Start(ctx, "new-cache-driver")
	defer span.End()

	switch config.System.CacheDriver {
	case "memory":
		return cache.NewMemoryDriver()
	case "redis":
		return cache.NewRedisDriver(params.RedisClient), nil
	default:
		err := fmt.Errorf("unknown cache driver: %s", config.System.CacheDriver)
		logrus.WithContext(ctx).WithError(err).Fatal("failed to create cache driver")
		return nil, err
	}
}
