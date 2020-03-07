package handlers

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/dpgil/url-shortener/types"
)

// Decode takes a short URL and returns the long URL.
func Decode(db *dynamodb.DynamoDB, tableName string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// parse short url from request
		shortURL, err := parseURLArg("/d/", r.URL.String())
		if err != nil {
			http.Error(w, "error parsing url", http.StatusBadRequest)
			return
		}

		// check if the short url is in the database
		params := &dynamodb.GetItemInput{
			TableName: aws.String(tableName),
			Key: map[string]*dynamodb.AttributeValue{
				"shortLink": {
					S: aws.String(shortURL),
				},
			},
		}

		resp, err := db.GetItem(params)
		if err != nil {
			http.Error(w, fmt.Sprintf("short link %s not found in database", shortURL), http.StatusNotFound)
			return
		}

		// unmarshal the dynamodb attribute values into our struct
		var mapping types.Mapping
		err = dynamodbattribute.UnmarshalMap(resp.Item, &mapping)
		if err != nil {
			http.Error(w, "error unmarshaling response from dynamo", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, mapping.LongLink)
	}
}
