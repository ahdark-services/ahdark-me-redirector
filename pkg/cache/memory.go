package cache

import (
	"context"
	"strings"
	"time"

	"github.com/dgraph-io/ristretto"
	"github.com/samber/lo"
	"go.opentelemetry.io/otel/attribute"
)

type MemoryDriver struct {
	client *ristretto.Cache
}

func NewMemoryDriver() (Driver, error) {
	c, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	if err != nil {
		return nil, err
	}

	return &MemoryDriver{
		client: c,
	}, nil
}

// Get implements Driver.Get
func (d *MemoryDriver) Get(ctx context.Context, key string) (any, bool) {
	ctx, span := tracer.Start(ctx, "memory.get")
	defer span.End()

	if span.IsRecording() {
		span.SetAttributes(
			attribute.String("key", key),
		)
	}

	value, found := d.client.Get(key)
	if !found {
		return nil, false
	}

	return value, true
}

// GetWithTTL implements Driver.GetWithTTL
func (d *MemoryDriver) GetWithTTL(ctx context.Context, key string) (any, time.Duration, error) {
	ctx, span := tracer.Start(ctx, "memory.get-with-ttl")
	defer span.End()

	if span.IsRecording() {
		span.SetAttributes(
			attribute.String("key", key),
		)
	}

	value, found := d.client.Get(key)
	if !found {
		return nil, 0, nil
	}

	ttl, _ := d.client.GetTTL(key)

	return value, ttl, nil
}

// GetMulti implements Driver.GetMulti
func (d *MemoryDriver) GetMulti(ctx context.Context, keys []string) (map[string]any, []string, error) {
	ctx, span := tracer.Start(ctx, "memory.get-multi")
	defer span.End()

	if span.IsRecording() {
		span.SetAttributes(
			attribute.String("keys", strings.Join(keys, ", ")),
		)
	}

	values := make(map[string]any, len(keys))
	var misses []string

	for _, key := range keys {
		value, found := d.client.Get(key)
		if !found {
			misses = append(misses, key)
			continue
		}

		values[key] = value
	}

	return values, misses, nil
}

// Set implements Driver.Set
func (d *MemoryDriver) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	ctx, span := tracer.Start(ctx, "memory.set")
	defer span.End()

	if span.IsRecording() {
		span.SetAttributes(
			attribute.String("key", key),
			attribute.String("ttl", ttl.String()),
		)
	}

	if ttl < 0 {
		ttl = 0 // ristretto does not support negative TTL
	}

	d.client.SetWithTTL(key, value, 1, ttl)

	return nil
}

// SetMulti implements Driver.SetMulti
func (d *MemoryDriver) SetMulti(ctx context.Context, values map[string]any, prefix string) error {
	ctx, span := tracer.Start(ctx, "memory.set-multi")
	defer span.End()

	keys := lo.MapToSlice(values, func(key string, value any) string {
		return prefix + key
	})

	if span.IsRecording() {
		span.SetAttributes(
			attribute.String("keys", strings.Join(keys, ", ")),
			attribute.String("prefix", prefix),
		)
	}

	for _, key := range keys {
		d.client.Set(key, values[key], 1)
	}

	return nil
}

// Delete implements Driver.Delete
func (d *MemoryDriver) Delete(ctx context.Context, key string) error {
	ctx, span := tracer.Start(ctx, "memory.delete")
	defer span.End()

	if span.IsRecording() {
		span.SetAttributes(
			attribute.String("key", key),
		)
	}

	d.client.Del(key)

	return nil
}

// DeleteMulti implements Driver.DeleteMulti
func (d *MemoryDriver) DeleteMulti(ctx context.Context, keys []string) error {
	ctx, span := tracer.Start(ctx, "memory.delete-multi")
	defer span.End()

	if span.IsRecording() {
		span.SetAttributes(
			attribute.String("keys", strings.Join(keys, ", ")),
		)
	}

	for _, key := range keys {
		d.client.Del(key)
	}

	return nil
}
