package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // parse arguments, you have to call this by yourself

	// todo: delete this block
	fmt.Println(r.Form) // print form information in server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	fmt.Fprintf(w, "redirect") // send data to client side
}

func encode(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "encode")
}

func decode(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "decode")
}

func main() {
	http.HandleFunc("/decode/", decode)
	http.HandleFunc("/encode/", encode)
	http.HandleFunc("/", redirect)
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
