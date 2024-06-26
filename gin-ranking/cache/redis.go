package cache

import (
	"context"
	"gin-ranking/config"

	"github.com/redis/go-redis/v9"
)

var (
	Rdb  *redis.Client
	Rctx context.Context
)

func init() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.RedisAddress,
		Password: config.RedisPassword,
		DB:       config.RedisDb,
	})
	Rctx = context.Background()
}

func Zscore(id int, score float64) redis.Z {
	return redis.Z{
		Score:  score,
		Member: id,
	}
}
