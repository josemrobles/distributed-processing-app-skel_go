package main

import (
	"github.com/gorilla/mux"
	"github.com/streadway/amqp"
	"log"
	"net/http"
	"os"
)

var (
	rmqConn      *amqp.Connection = nil
	payloadQueue                  = make(chan string)
	rmqServer                     = "amqp://guest:guest@" + os.Getenv("MESSAGEQUEUESERVER_PORT_5672_TCP_ADDR") + ":" + os.Getenv("MESSAGEQUEUESERVER_PORT_5672_TCP_PORT")
)

func init() {
	rmqConn = connect(rmqServer)

	for i := 0; i < 4; i++ {
		go payloadProcessor(payloadQueue)
	}

}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/v1/", doWork).Methods("POST")
	log.Fatal(http.ListenAndServe(":1337", router))
}
