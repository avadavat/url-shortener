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

// Redirect decodes a link and redirects the client.
func Redirect(db dynamodbiface.DynamoDBAPI, tableName string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// parse short link from request
		shortLink, err := util.ParseURLArg("/d/", r.URL.String())
		if err != nil {
			http.Error(w, "error parsing url", http.StatusBadRequest)
			return
		}

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
			http.Error(w, fmt.Sprintf("short link %s not found in database", shortLink), http.StatusNotFound)
			return
		}

		// unmarshal the dynamodb attribute values into our struct
		var mapping types.Mapping
		err = dynamodbattribute.UnmarshalMap(resp.Item, &mapping)
		if err != nil {
			http.Error(w, "error unmarshaling response from dynamo", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, mapping.LongLink, http.StatusPermanentRedirect)
	}
}
