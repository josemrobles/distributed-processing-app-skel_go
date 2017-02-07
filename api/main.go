package main

import (
	"github.com/gorilla/mux"
	"github.com/josemrobles/conejo"
	"log"
	"net/http"
	"os"
)

var (
	rmq      = conejo.Connect("amqp://guest:guest@localhost:5672")
	queue    = conejo.Queue{Name: "someQueue", Durable: false, Delete: false, Exclusive: false, NoWait: false}
	exchange = conejo.Exchange{Name: "someExchange", Type: "topic", Durable: true, AutoDeleted: false, Internal: false, NoWait: false}
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/v1/", doWork).Methods("POST")
	log.Fatal(http.ListenAndServe(":80", router))
}
