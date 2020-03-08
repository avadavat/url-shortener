package handlers

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/dpgil/url-shortener/types"
	"github.com/dpgil/url-shortener/util"
)

// Encode takes a long link, then generates, stores, and returns a short link.
func Encode(db dynamodbiface.DynamoDBAPI, tableName string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// parse long link from request
		longLink, err := util.ParseURLArg("/e/", r.URL.String())
		if err != nil {
			http.Error(w, "error parsing url", http.StatusBadRequest)
			return
		}

		var shortLink string
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

			_, err := db.GetItem(params)
			if err != nil {
				// The shortlink does not already exist in the database.
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
