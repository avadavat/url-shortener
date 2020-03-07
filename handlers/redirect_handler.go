package handlers

import (
	"fmt"
	"net/http"
)

// Redirect decodes a URL and redirects the client.
func Redirect(w http.ResponseWriter, r *http.Request) {
	// parse short url from request
	shortURL, err := parseURLArg("/r/", r.URL.String())
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

	http.Redirect(w, r, longURL, http.StatusPermanentRedirect)
}
