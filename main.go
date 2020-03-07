package main

import (
	"log"
	"net/http"

	"github.com/dpgil/url-shortener/handlers"
)

func main() {
	http.HandleFunc("/d/", handlers.Decode)
	http.HandleFunc("/e/", handlers.Encode)
	http.HandleFunc("/r/", handlers.Redirect)
	// todo: is 9090 the right port?
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
