package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hiamthach108/dreon-backend-service/config"
	"github.com/hiamthach108/dreon-sdk/logger"
	"github.com/redis/go-redis/v9"
)

const (
	ErrCacheNil = redis.Nil
)

type appCache struct {
	serviceName string
	logger      logger.ILogger
	redisClient *redis.Client
}

func NewAppCache(config *config.AppConfig, logger logger.ILogger) (ICache, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.Cache.RedisHost + ":" + config.Cache.RedisPort,
		Password: config.Cache.RedisPassword,
		DB:       config.Cache.RedisDB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := redisClient.Ping(ctx).Err(); err != nil {
		logger.Error("Failed to connect to Redis", "error", err)
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	logger.Info("Connected to Redis successfully")

	return &appCache{
		serviceName: config.App.Name,
		logger:      logger,
		redisClient: redisClient,
	}, nil
}

func (c *appCache) SetNX(ctx context.Context, key string, ttl time.Duration) (bool, error) {
	return c.redisClient.SetNX(ctx, c.prefixedKey(key), "1", ttl).Result()
}

func (c *appCache) Set(key string, value any, expireTime *time.Duration) error {
	var data any
	switch v := value.(type) {
	case string, int, int64, float64, bool:
		data = v
	default:
		jsonData, err := json.Marshal(value)
		if err != nil {
			return fmt.Errorf("failed to marshal value: %w", err)
		}
		data = jsonData
	}

	return c.redisClient.Set(context.Background(), c.prefixedKey(key), data, *expireTime).Err()
}

func (c *appCache) Get(key string, data any) error {
	val, err := c.redisClient.Get(context.Background(), c.prefixedKey(key)).Result()
	if err != nil {
		return err
	}
	if err := json.Unmarshal([]byte(val), data); err != nil {
		return err
	}
	return nil
}

func (c *appCache) Delete(key string) error {
	return c.redisClient.Del(context.Background(), c.prefixedKey(key)).Err()
}

func (c *appCache) Clear() error {
	return c.redisClient.FlushAll(context.Background()).Err()
}

func (c *appCache) ClearWithPrefix(prefix string) error {
	ctx := context.Background()
	keys, err := c.redisClient.Keys(ctx, c.prefixedKey(fmt.Sprintf("%s*", prefix))).Result()
	if err != nil {
		return err
	}
	if len(keys) > 0 {
		return c.redisClient.Del(ctx, keys...).Err()
	}
	return nil
}

func (c *appCache) prefixedKey(key string) string {
	if c.serviceName == "" {
		return key
	}
	return c.serviceName + ":" + key
}
