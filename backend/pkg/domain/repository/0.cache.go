package repository

import (
	"context"
	"fmt"
	"github.com/cocoide/tech-guide/pkg/usecase/parser"
	"time"

	"github.com/cocoide/tech-guide/pkg/domain/model/dto"
)

//go:generate mockgen -source=0.cache.go -destination=../../../mock/repository/cache.go
type CacheRepo interface {
	WithCtx(ctx context.Context) CacheRepo
	Get(key string) (string, bool, error)
	Set(key string, value string, expire time.Duration) error
	Delete(key string) error
	Update(key string, value string, expire time.Duration) error
	ExtendExpiry(key string, extension time.Duration) error
	AddSortedSet(key string, member interface{}, score float64, expire time.Duration) error
	GetAllSortedSet(key string) ([]dto.SortedSet, error)
	GetHashField(key string, fieldKey string) (string, error)
	SetHashField(key string, field map[string]interface{}, expire time.Duration) error
	GetAllHashFields(key string) (map[string]string, error)
}

func GetOrSet[T any](
	cache CacheRepo,
	key string, expire time.Duration,
	callback func(values ...interface{}) (T, error),
) (T, error) {
	var result T
	str, exist, err := cache.Get(key)
	if err != nil {
		return result, fmt.Errorf("Failed to cache")
	}
	if exist {
		result, err := parser.Deserialize[T](str)
		if exist {
			return result, err
		}
		return result, nil
	}
	// 以下、Cacheが存在しない場合
	result, err = callback()
	if err != nil {
		return result, err
	}
	str, err = parser.Serialize(result)
	if err != nil {
		return result, err
	}
	if err := cache.Set(key, str, expire); err != nil {
		return result, err
	}
	return result, nil
}
