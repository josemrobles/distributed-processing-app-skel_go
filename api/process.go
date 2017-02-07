package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func doWork(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("ERROR: Could not read POST data - %s", err)
		w.WriteHeader(500)
	} else {
		err := conejo.Publish(rmq, queue, exchange, string([]byte(b)))
		if err != nil {
			log.Printf("ERROR: Could not publish message %s", err)
		} else {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, "SOME SORT OF CONFIRMATIONN HERE")
		}
	}
}
