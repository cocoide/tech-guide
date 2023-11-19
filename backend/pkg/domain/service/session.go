package service

import (
	"github.com/google/uuid"
	"time"
)

type SessionService[T any] interface {
	Get(key uuid.UUID) (T, error)
	Del(key uuid.UUID) error
	New(value T, expire time.Duration) (uuid.UUID, error)
	Set(key uuid.UUID, value T, expire time.Duration) error
}
