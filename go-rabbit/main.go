package main

import (
	"fmt"
	"sync"

	"github.com/streadway/amqp"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go initProducer()
	go initConsumer()
	wg.Wait()
}

func initConsumer() {
	defer wg.Done()
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume("TestQueue", "", true, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("received Message %s", d.Body)
		}
	}()

	fmt.Println("success")
	<-forever
}

func initProducer() {
	defer wg.Done()
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Successfully Connected To our RabbitMQ")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("TestQueue", false, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(q)

	err = ch.Publish("", "TestQueue", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("hello word"),
	})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Successfully")

}
