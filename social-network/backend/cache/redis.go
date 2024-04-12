package cache

import (
	"context"
	"social-network/config"
	"social-network/middleware/logger"

	"github.com/redis/go-redis/v9"
)

var (
	Rdb  *redis.Client
	Rctx context.Context
)

func init() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.RedisHost,
		Password: config.RedisPassword,
		DB:       config.RedisDB,
	})
	Rctx = context.Background()

	_, err := Rdb.Ping(Rctx).Result()
	if err != nil {
		logger.Error(map[string]interface{}{"Redis connect error": err.Error()})
	}

}
