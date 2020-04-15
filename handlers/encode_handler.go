package handlers

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/dpgil/url-shortener/types"
	"github.com/dpgil/url-shortener/util"
)

const maxRetries = 3

// Encode takes a long link, then generates, stores, and returns a short link.
func Encode(db dynamodbiface.DynamoDBAPI, tableName string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// parse long link from request
		q := r.URL.Query()
		longLink := q.Get("q")
		longLink, err := url.QueryUnescape(longLink)
		if err != nil {
			http.Error(w, "error with QueryUnescape on longLink param", http.StatusBadRequest)
			return
		}
		if longLink == "" {
			http.Error(w, "missing query param", http.StatusBadRequest)
			return
		}

		var shortLink string
		retries := 0
		for {
			// Keep generating a short link until we find one that doesn't already exist.
			shortLink = util.GenerateShortLink()

			// check if the short link is in the database
			params := &dynamodb.GetItemInput{
				TableName: aws.String(tableName),
				Key: map[string]*dynamodb.AttributeValue{
					"shortLink": {
						S: aws.String(shortLink),
					},
				},
			}

			resp, err := db.GetItem(params)
			if err != nil {
				http.Error(w, "dynamodb error", http.StatusInternalServerError)
			}

			if len(resp.Item) == 0 {
				// The shortlink does not already exist in the database.
				break
			}

			retries++
			if retries >= maxRetries {
				http.Error(w, "error generating unique short link", http.StatusInternalServerError)
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

		params := &dynamodb.PutItemInput{
			TableName: aws.String(tableName),
			Item:      dbmapping,
		}

		_, err = db.PutItem(params)
		if err != nil {
			http.Error(w, fmt.Sprintf("error storing in dynamo: %s", err.Error()), http.StatusInternalServerError)
			return
		}

		// return the short link
		fmt.Fprintf(w, shortLink)
	}
}
