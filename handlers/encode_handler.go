package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/dpgil/url-shortener/types"
)

const shortURLLength = 4

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// todo: replace with db
var urlMapping = make(map[string]string)

// Encode takes a long URL, then generates, stores, and returns a short URL.
func Encode(db *dynamodb.DynamoDB, tableName string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// parse long url from request
		longURL, err := parseURLArg("/e/", r.URL.String())
		if err != nil {
			http.Error(w, "error parsing url", http.StatusBadRequest)
			return
		}

		// generate short url
		var shortURL string
		for {
			// Keep generating a short url until we find one that doesn't already exist.
			shortURL = generateShortURL()
			// todo: check dynamo instead of the internal map
			if _, ok := urlMapping[shortURL]; !ok {
				break
			}
		}

		// store the new mapping
		mapping := types.Mapping{
			ShortLink: shortURL,
			LongLink:  longURL,
		}

		dbmapping, err := dynamodbattribute.MarshalMap(mapping)
		if err != nil {
			http.Error(w, "error marshaling dynamo mapping", http.StatusInternalServerError)
			return
		}

		// todo: delete this once dynamo works
		// todo: allow table name to be configured
		params := &dynamodb.PutItemInput{
			TableName: aws.String(tableName),
			Item:      dbmapping,
		}

		_, err = db.PutItem(params)
		if err != nil {
			http.Error(w, fmt.Sprintf("error storing in dynamo: %s", err.Error()), http.StatusInternalServerError)
			return
		}

		urlMapping[shortURL] = longURL

		// return the short url
		fmt.Fprintf(w, shortURL)
	}
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
