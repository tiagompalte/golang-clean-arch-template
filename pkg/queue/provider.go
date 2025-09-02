package queue

import "github.com/tiagompalte/golang-clean-arch-template/configs"

func ProviderRabbitMqSet(config configs.Config) Queue[RabbitMqCommand] {
	return NewRabbitMq(config.Queue.ConnectionSource)
}
