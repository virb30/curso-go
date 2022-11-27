package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	ID  int64
	Msg string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)

	var i int64 = 0
	// RabbitMQ
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			time.Sleep(time.Second * 2)
			msg := Message{i, "Hello from RabbitMQ"}
			c1 <- msg
		}
	}()

	// Kafka
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			time.Sleep(time.Second * 1)
			msg := Message{i, "Hello from Kafka"}
			c2 <- msg
		}
	}()

	// for i := 0; i < 3; i++ {
	for {
		select {
		case msg := <-c1:
			fmt.Printf("received from rabbitmq: ID: %d - %s\n", msg.ID, msg.Msg)
		case msg := <-c2:
			fmt.Printf("received from kafka: ID: %d - %s\n", msg.ID, msg.Msg)
		case <-time.After(time.Second * 3):
			println("timeout")
			// default:
			// 	println("default")
		}
	}
}
