package handlers

import (
	"fmt"
	"net/http"
)

// Decode takes a short URL and returns the long URL.
func Decode(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "decode")
}
