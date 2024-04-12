package config

// 默认的 RabbitMQ 连接配置
const (
	RabbitMQURLDefault      = "amqp://root:123456@localhost:5672/"
	RabbitMQUsernameDefault = "root"
	RabbitMQPasswordDefault = "123456"
)

var (
	RabbitMQURL      string
	RabbitMQUsername string
	RabbitMQPassword string
)

func init() {
	RabbitMQURL = getEnv("RABBITMQ_URL", RabbitMQURLDefault)
	RabbitMQUsername = getEnv("RABBITMQ_DEFAULT_USER", RabbitMQUsernameDefault)
	RabbitMQPassword = getEnv("RABBITMQ_DEFAULT_PASS", RabbitMQPasswordDefault)
}
