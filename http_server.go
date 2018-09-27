package main

import (
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	address := flag.String("address", "0.0.0.0", "Listen address")
	port := flag.String("port", "18081", "Listen port")
	flag.Parse()

	addressAndPort := fmt.Sprintf("%s:%s", *address, *port)
	log.Println("Starting HTTP server at >>>>> ", addressAndPort)

	http.HandleFunc("/spycommand/reset", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" {
			// reset requestsHub
			for key, value := range r.URL.Body() {
				fmt.Println("\n param name:", key, "param value:", value)
			}
		}
	})
	http.HandleFunc("/spycommand/verify", func(w http.ResponseWriter, r *http.Request) {
		// Parameters -->  path and  method
		if r.Method == "GET" {
			// find from requestsHub
			for key, value := range r.URL.Query() {
				fmt.Println("\n param name:", key, "param value:", value)
			}
		}
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf(">>>>>> hello path: %q method: %q ", html.EscapeString(r.URL.Path), r.Method)
		for key, value := range r.URL.Query() {
			fmt.Println("\n k:", key, "v:", value)
		}
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe(addressAndPort, nil))
}
