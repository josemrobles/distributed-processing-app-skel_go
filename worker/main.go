package main

import (
	"github.com/josemrobles/conejo"
)

var (
	rmq       = conejo.Connect("amqp://guest:guest@localhost:5672")
	queue     = conejo.Queue{Name: "someQueue", Durable: false, Delete: false, Exclusive: false, NoWait: false}
	exchange  = conejo.Exchange{Name: "someExchange", Type: "topic", Durable: true, AutoDeleted: false, Internal: false, NoWait: false}
	workQueue = make(chan string)
)

func init() {
	for i := 0; i < 4; i++ {
		go asyncProcessor(workQueue)
	}
}

func main() {
	err := conejo.Consume(rmq, queue, exchange, "W1", workQueue)
	if err != nil {
		print("ERROR: %q", err)
	}
}
