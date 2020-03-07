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
		// todo: error
		fmt.Println("redirect: error parsing short URL from request URL")
		return
	}

	// check if the short URL is in the mapping
	longURL, ok := urlMapping[shortURL]
	if !ok {
		// todo: error, DNE
		fmt.Println("short url " + shortURL + " does not exist in the db")
		return
	}

	fmt.Println("redirecting to: " + longURL)
	http.Redirect(w, r, longURL, http.StatusPermanentRedirect)
}
