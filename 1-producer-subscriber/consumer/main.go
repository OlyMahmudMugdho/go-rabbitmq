package main

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	// Consumer = the receiver.
	conn, _ := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	ch, _ := conn.Channel()

	// Consume reads messages from the queue.
	msgs, _ := ch.Consume("basic_queue", "", true, false, false, false, nil)

	log.Println("waiting for messages...")
	for msg := range msgs {
		log.Println("received:", string(msg.Body))
	}
}

