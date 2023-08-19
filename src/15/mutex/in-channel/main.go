package main

import (
	"fmt"
	"log"
	"net/http"
)

var counter = make(chan int)

func main() {
	go generator()

	http.HandleFunc("/", handler)

	fmt.Println("listening on :9000")
	log.Fatal(http.ListenAndServe(":9000", nil)) // nolint
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%s] %s", r.Method, r.URL.String())
	fmt.Fprintf(w, "number %d", <-counter)
}

func generator() {
	for i := 0; ; i++ {
		counter <- i
	}
}
