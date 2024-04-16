package queue

import (
	"encoding/json"
	"social-network/cache"
	"social-network/config"
	"social-network/middleware/logger"
	"social-network/models"

	"github.com/streadway/amqp"
)

var (
	RabbitMQConn *amqp.Connection
	err          error
)

func init() {
	RabbitMQConn, err = amqp.Dial(config.RabbitMQURL)
	if err != nil {
		logger.Error(map[string]interface{}{"rabbitmq connect error": err.Error()})
		panic("Failed to connect to RabbitMQ, terminating application.")
	}
}
func getChannelDefualt() (*amqp.Channel, error) {
	ch, err := RabbitMQConn.Channel()
	if err != nil {
		logger.Error(map[string]interface{}{"error": "Failed to open a channel", "detail": err.Error()})
		return nil, err
	}
	return ch, nil
}

func declareQueueDeafult(ch *amqp.Channel, queueName string) (amqp.Queue, error) {
	q, err := ch.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		logger.Error(map[string]interface{}{"Connect RabbitMQ queue:": err.Error})
	}
	return q, err
}

func publishMessageDefualt(ch *amqp.Channel, queueName string, body []byte) error {
	err := ch.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		logger.Error(map[string]interface{}{"error": "Failed to publish message", "queue": queueName, "detail": err.Error()})
	}
	return err
}

func consumeMessageDefualt(ch *amqp.Channel, queueName string) (<-chan amqp.Delivery, error) {
	msgs, err := ch.Consume(
		queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logger.Error(map[string]interface{}{"error": "Failed to publish message", "queue": queueName, "detail": err.Error()})
	}
	return msgs, err
}

const (
	QAction_ToSQL   = "LikeAciton_ToSQL"
	QDelKey_ToRedis = "DelKey_ToRedis"
)

type LikeMsg struct {
	UserID   uint64 `json:"user_id"`   // 用户ID
	MomentID uint64 `json:"moment_id"` // 朋友圈消息ID
	Action   bool   `json:"action"`
}

type DelMsg struct {
	MomentID uint64 `json:"moment_id"`
}

func PublishMsg(msg interface{}, QName string) error {
	ch, err := getChannelDefualt()
	if err != nil {
		return err
	}
	defer ch.Close()
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return publishMessageDefualt(ch, QName, body)
}

func RedisConsumeDelMsg() {
	ch, err := getChannelDefualt()
	if err != nil {
		return
	}
	defer ch.Close()
	q, err := declareQueueDeafult(ch, QDelKey_ToRedis)
	if err != nil {
		return
	}
	msgs, err := consumeMessageDefualt(ch, q.Name)
	if err != nil {
		return
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			var msg DelMsg
			err := json.Unmarshal(d.Body, &msg)
			if err != nil {
				logger.Error(map[string]interface{}{"RabbitMQ Unmarshal error": err.Error()})
				d.Nack(false, true)
				continue
			}
			if err = cache.DelMonent(msg.MomentID); err != nil {
				logger.Error(map[string]interface{}{"RabbitMQ RedisConsumeDelMsg error": err.Error()})
				d.Nack(false, true)
				continue
			}
			d.Ack(false)
		}
	}()
	<-forever
}

func SQLConsumeLikeMsg() {
	ch, err := getChannelDefualt()
	if err != nil {
		return
	}
	defer ch.Close()
	q, err := declareQueueDeafult(ch, QAction_ToSQL)
	if err != nil {
		return
	}
	msgs, err := consumeMessageDefualt(ch, q.Name)
	if err != nil {
		return
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			var action LikeMsg
			err := json.Unmarshal(d.Body, &action)
			if err != nil {
				logger.Error(map[string]interface{}{"RabbitMQ Unmarshal error": err.Error()})
				d.Nack(false, true)
				continue
			}
			if action.Action {
				_, err = models.AddLike(action.UserID, action.MomentID)
			} else {
				err = models.DelLike(action.UserID, action.MomentID)
			}
			if err != nil {
				logger.Error(map[string]interface{}{"RabbitMQ ConsumeLikeActions error": err.Error()})
				d.Nack(false, true)
				continue
			}
			d.Ack(false)
		}
	}()
	<-forever
}

func SyncData() {
	go RedisConsumeDelMsg() //mysql接收redis数据
	go SQLConsumeLikeMsg()  //redis接收mysql数据
}
