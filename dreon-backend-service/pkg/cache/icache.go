package cache

import (
	"context"
	"time"
)

type ICache interface {
	SetNX(ctx context.Context, key string, ttl time.Duration) (bool, error)
	Set(key string, value any, expireTime *time.Duration) error
	Get(key string, data any) error
	Delete(key string) error
	Clear() error
	ClearWithPrefix(prefix string) error
}
