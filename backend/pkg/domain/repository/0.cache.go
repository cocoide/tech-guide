package repository

import (
	"context"
	"time"

	"github.com/cocoide/tech-guide/pkg/domain/model/dto"
)

//go:generate mockgen -source=0.cache.go -destination=../../../mock/repository/cache.go
type CacheRepo interface {
	WithCtx(ctx context.Context) CacheRepo
	Get(key string) (string, error)
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