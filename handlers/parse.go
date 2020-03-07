package handlers

import "errors"

// Takes a full url trims off the handler.
// i.e. parseURLArg("/handler/", "/handler/some/arg") => "some/arg"
func parseURLArg(handler, url string) (string, error) {
	l := len(handler)
	if len(url) < len(handler) {
		return "", errors.New("invalid input")
	}

	return url[l:], nil
}
