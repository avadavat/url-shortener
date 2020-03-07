package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

const shortURLLength = 5

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// todo: replace with db
var urlMapping = make(map[string]string)

// Encode takes a long URL, then generates, stores, and returns a short URL.
func Encode(w http.ResponseWriter, r *http.Request) {
	// parse long url from request
	longURL, err := parseURLArg("/e/", r.URL.String())
	if err != nil {
		// todo: error
		return
	}

	// generate short url
	var shortURL string
	for {
		shortURL = generateShortURL()
		if _, ok := urlMapping[shortURL]; !ok {
			urlMapping[shortURL] = longURL
			break
		}
	}

	// store the mapping
	urlMapping[shortURL] = longURL

	// return the short url
	fmt.Fprintf(w, shortURL)
}

// Generates a short URL
func generateShortURL() string {
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, shortURLLength)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
