package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

type foo struct{}

func (h *foo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from fooHandler, %q", html.EscapeString(r.URL.Path))
}

func main() {
	http.Handle("/foo", &foo{})

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from bar, %q", html.EscapeString(r.URL.Path))
	})

	log.Println("listening on port 9090")
	log.Fatal(http.ListenAndServe(":9090", nil)) // nolint
}
