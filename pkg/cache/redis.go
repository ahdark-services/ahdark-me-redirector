package cache

import (
	"bytes"
	"context"
	"encoding/gob"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisDriver struct {
	client *redis.Client
}

type redisObject struct {
	Value any
}

func init() {
	gob.Register(redisObject{})
}

func serialize(data any) ([]byte, error) {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	storeValue := redisObject{
		Value: data,
	}

	if err := enc.Encode(storeValue); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func deserialize(gobCode []byte) (any, error) {
	var obj redisObject
	buffer := bytes.NewBuffer(gobCode)
	dec := gob.NewDecoder(buffer)
	if err := dec.Decode(&obj); err != nil {
		return nil, err
	}

	return obj.Value, nil
}

func NewRedisDriver(client *redis.Client) Driver {
	return &RedisDriver{client: client}
}

// Get implements Driver.Get
func (d *RedisDriver) Get(ctx context.Context, key string) (any, bool) {
	ctx, span := tracer.Start(ctx, "RedisDriver.Get")
	defer span.End()

	result, err := d.client.Get(ctx, key).Bytes()
	if err != nil {
		span.RecordError(err)
		return nil, false
	}

	obj, err := deserialize(result)
	if err != nil {
		span.RecordError(err)
		return nil, false
	}

	return obj, true
}

// GetWithTTL implements Driver.GetWithTTL
func (d *RedisDriver) GetWithTTL(ctx context.Context, key string) (any, time.Duration, error) {
	ctx, span := tracer.Start(ctx, "RedisDriver.GetWithTTL")
	defer span.End()

	result, err := d.client.Get(ctx, key).Bytes()
	if err != nil {
		span.RecordError(err)
		return nil, 0, err
	}

	obj, err := deserialize(result)
	if err != nil {
		span.RecordError(err)
		return nil, 0, err
	}

	ttl, err := d.client.TTL(ctx, key).Result()
	if err != nil {
		span.RecordError(err)
		return nil, 0, err
	}

	return obj, ttl, nil
}

// GetMulti implements Driver.GetMulti
func (d *RedisDriver) GetMulti(ctx context.Context, keys []string) (map[string]any, []string, error) {
	ctx, span := tracer.Start(ctx, "RedisDriver.Gets")
	defer span.End()

	result, err := d.client.MGet(ctx, keys...).Result()
	if err != nil {
		span.RecordError(err)
		return nil, nil, err
	}

	res := make(map[string]any)
	missed := make([]string, 0, len(keys))
	for i, value := range result {
		if value == nil {
			missed = append(missed, keys[i])
			continue
		}

		decoded, err := deserialize([]byte(value.(string)))
		if err != nil || decoded == nil {
			missed = append(missed, keys[i])
		} else {
			res[keys[i]] = decoded
		}
	}

	return res, missed, nil
}

// Set implements Driver.Set
func (d *RedisDriver) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	ctx, span := tracer.Start(ctx, "RedisDriver.Set")
	defer span.End()

	b, err := serialize(value)
	if err != nil {
		span.RecordError(err)
		return err
	}

	if _, err := d.client.Set(ctx, key, b, ttl).Result(); err != nil {
		span.RecordError(err)
		return err
	}

	return nil
}

// SetMulti implements Driver.SetMulti
func (d *RedisDriver) SetMulti(ctx context.Context, values map[string]any, prefix string) error {
	ctx, span := tracer.Start(ctx, "RedisDriver.Sets")
	defer span.End()

	var slice []any
	for key, value := range values {
		b, err := serialize(value)
		if err != nil {
			span.RecordError(err)
			return err
		}

		slice = append(slice, prefix+key, b)
	}

	if _, err := d.client.MSet(ctx, slice...).Result(); err != nil {
		span.RecordError(err)
		return err
	}

	return nil
}

// Delete implements Driver.Delete
func (d *RedisDriver) Delete(ctx context.Context, key string) error {
	ctx, span := tracer.Start(ctx, "RedisDriver.Delete")
	defer span.End()

	if _, err := d.client.Del(ctx, key).Result(); err != nil {
		span.RecordError(err)
		return err
	}

	return nil
}

// DeleteMulti implements Driver.DeleteMulti
func (d *RedisDriver) DeleteMulti(ctx context.Context, keys []string) error {
	ctx, span := tracer.Start(ctx, "RedisDriver.Deletes")
	defer span.End()

	if _, err := d.client.Del(ctx, keys...).Result(); err != nil {
		span.RecordError(err)
		return err
	}

	return nil
}
