package queue

import (
	"encoding/json"
	"social-network/cache"
	"social-network/config"
	"social-network/middleware/logger"
	"social-network/models"
	"strconv"

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

type LikeRequest struct {
	UserID   uint64 `json:"user_id"`   // 用户ID
	MomentID uint64 `json:"moment_id"` // 朋友圈消息ID
	Action   bool   `json:"action"`
}

func PublishLikeAction(userID, momentID uint64, action bool) error {

	ch, err := RabbitMQConn.Channel()

	if err != nil {
		return err
	}
	defer ch.Close()

	liks := LikeRequest{
		UserID:   userID,
		MomentID: momentID,
		Action:   action,
	}

	body, err := json.Marshal(liks)

	if err != nil {
		return err
	}

	err = ch.Publish(
		"",           // exchange
		"likesQueue", // routing key (queue name)
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	return err
}

func ConsumeLikeActions() {
	ch, err := RabbitMQConn.Channel()
	if err != nil {
		logger.Error(map[string]interface{}{"Connect RabbitMQ  queue": err})

	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"likes_queue", // queue name
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	logger.Error(map[string]interface{}{"Failed to declare a queue": err})

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {
		return
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var action LikeRequest
			err := json.Unmarshal(d.Body, &action)
			if err != nil {
				logger.Error(map[string]interface{}{"RabbitMQ Unmarshal error": err.Error()})
				d.Nack(false, true)
			}
			if action.Action {
				_, err = models.AddLike(action.UserID, action.MomentID)
			} else {
				_, err = models.DelLike(action.UserID, action.MomentID)
			}
			if err != nil {
				logger.Error(map[string]interface{}{"RabbitMQ Consume error": err.Error()})
				d.Nack(false, true)
			}
			d.Ack(false)
		}
	}()

	<-forever
}

type LikeCounts struct {
	MomentId uint64 `json:"moment_id"`
	Count    uint64 `json:"count"`
}

func PublishLikeCounts(momentid, count uint64) error {
	ch, err := RabbitMQConn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	data := LikeCounts{
		MomentId: momentid,
		Count:    count,
	}

	body, err := json.Marshal(data)

	if err != nil {
		return err
	}

	err = ch.Publish(
		"",
		"likecounts_queue",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	return err

}

func ConsumeLikeCounts() {
	ch, err := RabbitMQConn.Channel()
	if err != nil {
		logger.Error(map[string]interface{}{"Connect RabbitMQ queue:": err.Error})
		return
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"likecounts_queue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logger.Error(map[string]interface{}{"Declare likecounts queue:": err.Error})
		return
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		logger.Error(map[string]interface{}{"Consume likecounts queue:": err.Error})
		return
	}

	forever := make(chan struct{})

	go func() {
		for d := range msgs {
			var likecounts LikeCounts
			err := json.Unmarshal(d.Body, &likecounts)
			if err != nil {
				logger.Error(map[string]interface{}{"RabbitMQ unmarshal error:": err.Error()})
				d.Nack(false, true)
				continue
			}
			momentid := strconv.FormatUint(likecounts.MomentId, 10)

			_, err = cache.Rdb.HSet(cache.Rctx, "likes:count", momentid, likecounts.Count).Result()

			if err != nil {
				logger.Error(map[string]interface{}{"RabbitMQ likecounts to Redis fail:": err.Error()})
				d.Nack(false, true)
				continue
			}
			d.Ack(false)
		}
	}()
	<-forever

}
func SyncData() {

	go ConsumeLikeActions()

}
