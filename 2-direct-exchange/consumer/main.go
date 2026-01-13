package main

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, _ := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	ch, _ := conn.Channel()

	exchange := "direct_logs"

	/*
	   Queue will only receive messages whose routing key matches "info".
	*/
	q, _ := ch.QueueDeclare("direct_info_queue", false, false, false, false, nil)

	/*
	   Binding:
	   This rule connects the queue with the exchange.
	   It says: send messages with key "info" to this queue.
	*/
	ch.QueueBind(q.Name, "info", exchange, false, nil)

	msgs, _ := ch.Consume(q.Name, "", true, false, false, false, nil)

	log.Println("waiting for direct match...")
	for msg := range msgs {
		log.Println("received:", string(msg.Body))
	}
}
