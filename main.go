package main

import (
	"fmt"
	"log"
	"net/http"
)

func multiplyByTwo(chanIn chan int, chanOut chan int) {
	for {
		res := <-chanIn
		chanOut <- res * 2
	}
}

func generateHandler(chanIn, chanOut chan int) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		chanIn <- 2
		res := <-chanOut
		fmt.Println("Value is now:", res)
		fmt.Fprintf(w, `{"value": %d}`, res)
	}
}

func main() {
	channelIn, channelOut := make(chan int), make(chan int)

	go multiplyByTwo(channelIn, channelOut)

	http.HandleFunc("/data", generateHandler(channelIn, channelOut))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
