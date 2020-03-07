package handlers

import (
	"fmt"
	"net/http"
)

// Redirect decodes a URL and redirects the client.
func Redirect(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "redirect")
}
