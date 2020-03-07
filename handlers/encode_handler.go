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

const shortLinkLength = 4

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// todo: replace with db
var urlMapping = make(map[string]string)

// Encode takes a long link, then generates, stores, and returns a short link.
func Encode(db *dynamodb.DynamoDB, tableName string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// parse long link from request
		longLink, err := parseURLArg("/e/", r.URL.String())
		if err != nil {
			http.Error(w, "error parsing url", http.StatusBadRequest)
			return
		}

		// generate short link
		var shortLink string
		for {
			// Keep generating a short link until we find one that doesn't already exist.
			shortLink = generateShortLink()
			// todo: check dynamo instead of the internal map
			if _, ok := urlMapping[shortLink]; !ok {
				break
			}
		}

		// store the new mapping
		mapping := types.Mapping{
			ShortLink: shortLink,
			LongLink:  longLink,
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

		urlMapping[shortLink] = longLink

		// return the short link
		fmt.Fprintf(w, shortLink)
	}
}

// Generates a short link
func generateShortLink() string {
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, shortLinkLength)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
