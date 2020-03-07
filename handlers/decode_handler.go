package handlers

import (
	"fmt"
	"net/http"
)

// Decode takes a short URL and returns the long URL.
func Decode(w http.ResponseWriter, r *http.Request) {
	// parse short url from request
	shortURL, err := parseURLArg("/d/", r.URL.String())
	if err != nil {
		// todo: error
		return
	}

	// check if the short URL is in the mapping
	longURL, ok := urlMapping[shortURL]
	if !ok {
		// todo: error, DNE
		fmt.Printf("short url " + shortURL + " does not exist in the db")
		return
	}

	fmt.Fprintf(w, longURL)
}
