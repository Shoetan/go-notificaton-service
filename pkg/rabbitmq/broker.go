package rabbitmq

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Set up a publisher func that sends the message
//set up receiver or consume rabbitmq sends the message to


func PublishToQueue(conn *amqp.Connection) error  {

	ch, err := conn.Channel()

	if err != nil {
		return err
	}

	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Println("Could not create queue")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := "Hello Genuis"

	err = ch.PublishWithContext(ctx,
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(body),
		})
		
		return err
	
}

func ReceiveFromQueue(conn *amqp.Connection) {
	
	ch, err := conn.Channel()

	if err != nil {
		log.Println("Could not create channel")
	}

	queue, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Println("Could not create queue")
	}

	msgs, err := ch.Consume(
		queue.Name,
		"",
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    
	)

	if err != nil {
		log.Println("Could not consume message")
	}

	go func ()  {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()


}
