// config/redis.go

package config

const (
	RedisHost_DEFAULT = "localhost:39002"
)

var (
	RedisHost     string
	RedisPassword string
	RedisDB       int
)

func init() {
	RedisHost = getEnv("REDIS_HOST", RedisHost_DEFAULT)
}
