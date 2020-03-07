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
		http.Error(w, "error parsing url", http.StatusBadRequest)
		return
	}

	// check if the short URL is in the mapping
	longURL, ok := urlMapping[shortURL]
	if !ok {
		http.Error(w, fmt.Sprintf("short url %s does not exist", shortURL), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, longURL)
}
