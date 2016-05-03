package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/streadway/amqp"
	"log"
	"net/http"
)

var (
	port = flag.String("port", "8080", "API Port")
	uri = flag.String("uri", "", "AMQP URI")
	rmqConn *amqp.Connection = nil
)

func init() {
	flag.Parse()
	log.Printf("API listening on port :%s", *port)
	rmqConn = connect(*uri)
}

func main() {
	p := ":" + *port
	router := mux.NewRouter().StrictSlash(true)
	//router.HandleFunc("api/v1/app/endpoint", endpointFunc).Methods("POST")
	log.Fatal(http.ListenAndServe(p, router))
}

// Function used to connect to the RabbitMQ server specified in rmqConn
func connect(amqpURI string) (ch *amqp.Connection) {
	connection, err := amqp.Dial(amqpURI)
	if err != nil {
		log.Printf("ERROR: Could not connect to RabbitMQ server - %q", err)
		return nil
	} else {
		log.Printf("Connected to RabbitMQ server - %s", amqpURI)
		return connection
	}
}
