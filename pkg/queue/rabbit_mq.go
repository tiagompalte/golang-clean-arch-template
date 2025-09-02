package queue

import (
	"context"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type rabbitMq struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewRabbitMq(connectionUrl string) Queue[RabbitMqCommand] {
	conn, err := amqp.Dial(connectionUrl)
	if err != nil {
		log.Fatalf("error to connect in rabbitmq: %v", err)
	}

	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("error to create a new channel: %v", err)
	}

	return rabbitMq{
		conn:    conn,
		channel: channel,
	}
}

func (q rabbitMq) IsHealthy(ctx context.Context) (bool, error) {
	return !q.conn.IsClosed(), nil
}

func (q rabbitMq) Command() RabbitMqCommand {
	return RabbitMqCommand{
		CreateChannel: func(ctx context.Context) (*amqp.Channel, error) {
			return q.conn.Channel()
		},
		CreateExchange: func(ctx context.Context) error {
			return q.channel.ExchangeBind()
		},
		CreateQueue: func(ctx context.Context) (*amqp.Queue, error) {
		},
	}
}
