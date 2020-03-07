package handlers

import (
	"fmt"
	"net/http"
)

// Encode takes a long URL, then generates, stores, and returns a short URL.
func Encode(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// parse long url from request

	// generate short url

	// store the mapping

	// return the short url

	fmt.Fprintf(w, "encode")
}
