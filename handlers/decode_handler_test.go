package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dpgil/url-shortener/handlers"
)

func TestDecodeHandler(t *testing.T) {
	t.Parallel()

	mockSvc := &mockDynamoDBClient{}
	mockTableName := "mock-table-name"
	handler := http.HandlerFunc(handlers.Decode(mockSvc, mockTableName))

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

	t.Run("shortlink doesn't exist", func(t *testing.T) {
		t.Parallel()

		r, err := http.NewRequest("GET", "/e/someShortLink", nil)
		if err != nil {
			t.Fatal(err)
		}
		w := httptest.NewRecorder()

		handler.ServeHTTP(w, r)

		if status := w.Code; status != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusNotFound)
		}
	})

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		r, err := http.NewRequest("GET", "/e/getItemSuccess", nil)
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
