package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"spillane.farm/gowebtemplate/internal/system"
	"time"
)

var rdb *redis.Client

func Init(cfg system.Config) {
	logrus.Info("Initializing Redis")
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:6379", cfg.RedisHost),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func Set(ctx context.Context, key string, value string, expiration time.Duration) error {
	err := rdb.Set(ctx, key, value, expiration).Err()
	if err != nil {
		logrus.WithError(err).Error("failed to write to cache")
	}
	return err
}

func Get(ctx context.Context, key string) (string, error) {
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		logrus.WithField("key", key).WithError(err).Error("failed to fetch from cache")
	}
	return val, err
}
