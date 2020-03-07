package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// Decode takes a short URL and returns the long URL.
func Decode(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// parse short url from request
	shortURL, err := parseShortURLFromRequestURL(r.URL.String())
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

// Parses the short URL from the request URL.
// todo: there has to be a better way to structure our URLs to take
// advantage of Go features
func parseShortURLFromRequestURL(requestURL string) (string, error) {
	// Expected in the format /decode/<longURL>
	parsed := strings.Split(requestURL, "/decode/")
	if len(parsed) < 2 {
		return "", errors.New("unsuccessfully parsed url")
	}

	return parsed[1], nil
}
