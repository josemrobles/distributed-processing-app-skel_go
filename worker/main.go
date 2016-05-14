package main

import (
	"github.com/streadway/amqp"
	"log"
	"os"
)

var (
	rmqConn   *amqp.Connection = nil
	workQueue                  = make(chan string)
	rmqServer                  = "amqp://guest:guest@" + os.Getenv("RABBITMQ_PORT_5672_TCP_ADDR") + ":" + os.Getenv("RABBITMQ_PORT_5672_TCP_PORT")
)

func init() {
	rmqConn = connect(rmqServer)
	for i := 0; i < 4; i++ {
		go asyncProcessor(workQueue)
	}
}

func main() {
	err := consume(rmqConn, os.Getenv("MESSAGEQUEUESERVER_EXCHANGE"), os.Getenv("MESSAGEQUEUESERVER_QUEUE"), "W1", workQueue)
	if err != nil {
		log.Printf("ERROR: %q", err)
	}
}
