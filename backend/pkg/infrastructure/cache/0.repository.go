package cache

import (
	"context"
	"errors"
	"time"

	"github.com/cocoide/tech-guide/pkg/domain/model/dto"
	"github.com/cocoide/tech-guide/pkg/domain/repository"
	"github.com/redis/go-redis/v9"
)

type Repository struct {
	redis *redis.Client
	ctx   context.Context
}

func NewRepository(redis *redis.Client, ctx context.Context) *Repository {
	return &Repository{redis: redis, ctx: ctx}
}

func (r *Repository) WithCtx(ctx context.Context) repository.CacheRepo {
	return &Repository{
		redis: r.redis,
		ctx:   ctx,
	}
}

func (r *Repository) Get(key string) (string, bool, error) {
	value, err := r.redis.Get(r.ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", false, nil
		}
		return "", false, err
	}
	return value, true, nil
}

func (r *Repository) Set(key string, value string, expire time.Duration) error {
	err := r.redis.Set(r.ctx, key, value, expire).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Delete(key string) error {
	err := r.redis.Del(r.ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Update(key string, value string, expire time.Duration) error {
	return r.Set(key, value, expire)
}

func (r *Repository) ExtendExpiry(key string, extension time.Duration) error {
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

func (r *Repository) AddSortedSet(key string, member interface{}, score float64, expire time.Duration) error {
	err := r.redis.ZAdd(r.ctx, key, redis.Z{
		Score:  score,
		Member: member,
	}).Err()
	if err != nil {
		return err
	}
	err = r.redis.Expire(r.ctx, key, expire).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetAllSortedSet(key string) ([]dto.SortedSet, error) {
	var result []dto.SortedSet
	redisZ, err := r.redis.ZRangeWithScores(r.ctx, key, 0, -1).Result()
	if err != nil {
		return nil, err
	}
	for _, v := range redisZ {
		result = append(result, dto.SortedSet{
			Score:  v.Score,
			Member: v.Member.(string),
		})
	}
	return result, nil
}

func (r *Repository) GetHashField(key string, fieldKey string) (string, error) {
	return r.redis.HGet(r.ctx, key, fieldKey).Result()
}

func (r *Repository) SetHashField(key string, field map[string]interface{}, expire time.Duration) error {
	if err := r.redis.HSet(r.ctx, key, field).Err(); err != nil {
		return nil
	}
	if expire > 0 {
		err := r.redis.Expire(r.ctx, key, expire).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Repository) GetAllHashFields(key string) (map[string]string, error) {
	result, err := r.redis.HGetAll(r.ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return result, nil
}
