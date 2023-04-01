package cache

import (
	"context"
	"go.opentelemetry.io/otel"
	"time"
)

var tracer = otel.Tracer("internal.cache")

type Driver interface {
	Get(ctx context.Context, key string) (any, bool)
	GetWithTTL(ctx context.Context, key string) (any, time.Duration, error)
	GetMulti(ctx context.Context, keys []string) (map[string]any, []string, error)
	Set(ctx context.Context, key string, value any, ttl time.Duration) error
	SetMulti(ctx context.Context, values map[string]any, prefix string) error
	Delete(ctx context.Context, key string) error
	DeleteMulti(ctx context.Context, keys []string) error
}
