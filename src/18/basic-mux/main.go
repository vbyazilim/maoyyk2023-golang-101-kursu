package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

type apiHandler struct{}

func (apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the %q", html.EscapeString(r.URL.Path))
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/api/", apiHandler{})
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// The "/" pattern matches everything, so we need to check
		// that we're at the root here.
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}
		fmt.Fprintf(w, "Welcome to the home page!")
	})

	log.Println("listening on port 9090")
	log.Fatal(http.ListenAndServe(":9090", mux)) // nolint
}
