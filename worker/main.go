package main

import (
	"github.com/streadway/amqp"
	"log"
	"os"
)

var (
	rmqConn   *amqp.Connection = nil
	workQueue                  = make(chan string)
	rmqServer                  = "amqp://guest:guest@" + os.Getenv("RABBITMQ_SERVER") + ":5672"
)

func init() {
	rmqConn = connect(rmqServer)
	for i := 0; i < 4; i++ {
		go asyncProcessor(workQueue)
	}
}

func main() {
	err := consume(rmqConn, os.Getenv("RMQ_EXCHANGE"), os.Getenv("RMQ_QUEUE"), "W1", workQueue)
	if err != nil {
		log.Printf("ERROR: %q", err)
	}
}
