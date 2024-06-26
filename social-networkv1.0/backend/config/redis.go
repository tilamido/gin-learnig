// config/redis.go

package config

const (
	RedisHost_DEFAULT = "localhost:6379"
)

var (
	RedisHost string
)

func init() {
	RedisHost = getEnv("REDIS_HOST", RedisHost_DEFAULT)
}
