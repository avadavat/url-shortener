package handlers

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Redirect decodes a link and redirects the client.
func Redirect(db *dynamodb.DynamoDB, tableName string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// parse short link from request
		shortLink, err := parseURLArg("/r/", r.URL.String())
		if err != nil {
			http.Error(w, "error parsing link", http.StatusBadRequest)
			return
		}

		// check if the short link is in the mapping
		longLink, ok := urlMapping[shortLink]
		if !ok {
			http.Error(w, fmt.Sprintf("short link %s does not exist", shortLink), http.StatusBadRequest)
			return
		}

		http.Redirect(w, r, longLink, http.StatusPermanentRedirect)
	}
}
