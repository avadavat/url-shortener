package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/decode/", decode)
	http.HandleFunc("/encode/", encode)
	http.HandleFunc("/", redirect)

	// todo: is 9090 the right port?
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
