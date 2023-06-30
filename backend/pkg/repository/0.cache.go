package repository

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheRepo interface {
	WithCtx(ctx context.Context) CacheRepo
	Get(key string) (string, error)
	Set(key string, value string, expire time.Duration) error
	Delete(key string) error
	Update(key string, value string, expire time.Duration) error
	ExtendExpiry(key string, extension time.Duration) error
}

type cacheRepo struct {
	redis *redis.Client
	ctx   context.Context
}

func NewCacheRepo(redis *redis.Client) CacheRepo {
	return &cacheRepo{
		redis: redis,
		ctx:   context.Background(),
	}
}

func (r *cacheRepo) WithCtx(ctx context.Context) CacheRepo {
	return &cacheRepo{
		redis: r.redis,
		ctx:   ctx,
	}
}

func (r *cacheRepo) Get(key string) (string, error) {
	value, err := r.redis.Get(r.ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", nil
		}
		return "", err
	}
	return value, nil
}

func (r *cacheRepo) Set(key string, value string, expire time.Duration) error {
	err := r.redis.Set(r.ctx, key, value, expire).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *cacheRepo) Delete(key string) error {
	err := r.redis.Del(r.ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *cacheRepo) Update(key string, value string, expire time.Duration) error {
	return r.Set(key, value, expire)
}

func (r *cacheRepo) ExtendExpiry(key string, extension time.Duration) error {
	ttl := r.redis.TTL(r.ctx, key).Val()
	if ttl <= 0 {
		return errors.New("key does not exist or has no expiration")
	}

	newExpiration := ttl + extension
	if newExpiration > 0 {
		return r.redis.Expire(r.ctx, key, newExpiration).Err()
	}

	return nil
}
