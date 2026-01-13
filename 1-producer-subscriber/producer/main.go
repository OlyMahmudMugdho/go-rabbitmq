package main

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	// Producer = the sender.
	conn, _ := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	ch, _ := conn.Channel()

	/*
	   Queue: A queue stores messages until a consumer reads them.
	   It's like a waiting line. First in, first out.
	*/
	q, _ := ch.QueueDeclare("basic_queue", false, false, false, false, nil)

	/*
	   Here we publish directly to the queue through the default exchange "".
	   The default exchange sends the message to a queue with the same name
	   as the routing key.
	*/
	body := "Hello from the producer"
	ch.Publish("", q.Name, false, false, amqp091.Publishing{
		Body: []byte(body),
	})

	log.Println("sent:", body)
}

