package main

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, _ := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	ch, _ := conn.Channel()

	/*
	   Exchange:
	   Producers send messages to exchanges, not queues.
	   A direct exchange routes messages when the routing key
	   exactly matches the binding key.
	*/
	exchange := "direct_logs"
	ch.ExchangeDeclare(exchange, "direct", false, false, false, false, nil)

	/*
	   Routing key:
	   This label helps the exchange decide where to send the message.
	*/
	routingKey := "info"

	body := "This is an info message"
	ch.Publish(exchange, routingKey, false, false, amqp091.Publishing{
		Body: []byte(body),
	})

	log.Println("sent:", body)
}
