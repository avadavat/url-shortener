package handlers_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/dpgil/url-shortener/handlers"
)

// Define a mock struct to be used in your unit tests of myFunc.
type mockDynamoDBClient struct {
	dynamodbiface.DynamoDBAPI
}

// Encode generates new shortlinks until it finds one that doesn't exist in the database.
// An error is expected for Encode to be successful.
func (m *mockDynamoDBClient) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	shortLink := *input.Key["shortLink"].S
	longLink := "someLongLink"
	if shortLink == "getItemSuccess" {
		output := &dynamodb.GetItemOutput{
			Item: map[string]*dynamodb.AttributeValue{
				"shortLink": {S: &shortLink},
				"longLink":  {S: &longLink},
			},
		}
		return output, nil
	}

	return nil, errors.New("item does not exist in the database")
}
func (m *mockDynamoDBClient) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	if *input.Item["longLink"].S == "putErrorUrl" {
		return nil, errors.New("dynamo PutItem error")
	}

	return nil, nil
}

func TestEncodeHandler(t *testing.T) {
	t.Parallel()

	mockSvc := &mockDynamoDBClient{}
	mockTableName := "mock-table-name"
	handler := http.HandlerFunc(handlers.Encode(mockSvc, mockTableName))

	t.Run("invalid request format", func(t *testing.T) {
		t.Parallel()

		r, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		w := httptest.NewRecorder()

		handler.ServeHTTP(w, r)

		if status := w.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}
	})

	t.Run("dynamo put error", func(t *testing.T) {
		t.Parallel()

		r, err := http.NewRequest("GET", "/e/putErrorUrl", nil)
		if err != nil {
			t.Fatal(err)
		}
		w := httptest.NewRecorder()

		handler.ServeHTTP(w, r)

		if status := w.Code; status != http.StatusInternalServerError {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusInternalServerError)
		}
	})

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		r, err := http.NewRequest("GET", "/e/https://google.com", nil)
		if err != nil {
			t.Fatal(err)
		}
		w := httptest.NewRecorder()

		handler.ServeHTTP(w, r)

		if status := w.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	})
}
