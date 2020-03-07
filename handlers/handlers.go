package handlers

import (
	"fmt"
	"net/http"
)

// Maps short url to long url
var shortenerMap = make(map[string]string)

// Redirect decodes a URL and redirects the client.
func Redirect(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "redirect")
}

// Encode takes a long URL, then generates, stores, and returns a short URL.
func Encode(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// parse long url from request

	// generate short url

	// store the mapping

	// return the short url

	fmt.Fprintf(w, "encode")
}

// Decode takes a short URL and returns the long URL.
func Decode(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "decode")
}

// func redirect(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm() // parse arguments, you have to call this by yourself

// 	// todo: delete this block
// 	fmt.Println(r.Form) // print form information in server side
// 	fmt.Println("path", r.URL.Path)
// 	fmt.Println("scheme", r.URL.Scheme)
// 	fmt.Println(r.Form["url_long"])
// 	for k, v := range r.Form {
// 		fmt.Println("key:", k)
// 		fmt.Println("val:", strings.Join(v, ""))
// 	}

// 	fmt.Fprintf(w, "redirect") // send data to client side
// }
