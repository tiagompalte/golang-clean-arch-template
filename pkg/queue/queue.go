package queue

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/healthcheck"
)

type RabbitMqCommand struct {
	CreateChannel  func(ctx context.Context) (*amqp.Channel, error)
	CreateExchange func(ctx context.Context) error
	CreateQueue    func(ctx context.Context) (*amqp.Queue, error)
}

type Command interface {
	RabbitMqCommand
}

type Queue[T RabbitMqCommand] interface {
	healthcheck.HealthCheck
	Command() T
}
