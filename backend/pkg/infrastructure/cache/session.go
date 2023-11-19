package cache

import (
	"encoding/json"
	"fmt"
	"github.com/cocoide/tech-guide/pkg/domain/service"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"time"
)

func NewSessionService[T any](redis *redis.Client, ctx context.Context) service.SessionService[T] {
	return sessionService[T]{redis: redis, ctx: ctx}
}

type sessionService[T any] struct {
	redis *redis.Client
	ctx   context.Context
}

func (s sessionService[T]) Get(key uuid.UUID) (T, error) {
	var result T
	sessionKey := fmt.Sprintf("session.%s", key.String())
	b, err := s.redis.Get(s.ctx, sessionKey).Bytes()
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(b, &result); err != nil {
		return result, nil
	}
	return result, nil
}

func (s sessionService[T]) New(value T, expire time.Duration) (uuid.UUID, error) {
	u := uuid.New()
	b, err := json.Marshal(value)
	if err != nil {
		return u, err
	}
	sessionKey := fmt.Sprintf("session.%s", u.String())
	if err := s.redis.Set(s.ctx, sessionKey, b, expire).Err(); err != nil {
		return u, err
	}
	return u, nil
}

func (s sessionService[T]) Set(key uuid.UUID, value T, expire time.Duration) error {
	b, err := json.Marshal(value)
	if err != nil {
		return err
	}
	sessionKey := fmt.Sprintf("session.%s", key.String())
	if err := s.redis.Set(s.ctx, sessionKey, b, expire).Err(); err != nil {
		return err
	}
	return nil
}

func (s sessionService[T]) Del(key uuid.UUID) error {
	sessionKey := fmt.Sprintf("session.%s", key.String())
	return s.redis.Del(s.ctx, sessionKey).Err()
}
