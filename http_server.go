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
	port := flag.String("port", "18080", "Listen port")
	flag.Parse()

	addressAndPort := fmt.Sprintf("%s:%s", *address, *port)
	log.Println("[http2amqp] Starting HTTP server at ", addressAndPort)
	http.HandleFunc("/kk", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf(">>>>>> hello path: %q method: %q ", html.EscapeString(r.URL.Path), r.Method)
		for key, value := range r.URL.Query() {
			fmt.Println("\n k:", key, "v:", value)
		}
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe(addressAndPort, nil))
}
