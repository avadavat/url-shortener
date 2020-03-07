package main

import (
	"log"
	"net/http"

	"github.com/dpgil/url-shortener/handlers"
)

func main() {
	http.HandleFunc("/decode/", handlers.Decode)
	http.HandleFunc("/encode/", handlers.Encode)
	http.HandleFunc("/", handlers.Redirect)
	// todo: is 9090 the right port?
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
