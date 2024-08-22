package rabbitmq

import (
  amqp "github.com/rabbitmq/amqp091-go"
)

func RabbitMqConn() (*amqp.Connection, error) {
	 conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	 return conn, err
}