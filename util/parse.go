package util

import "errors"

// ParseURLArg takes a full url trims off the handler.
// i.e. ParseURLArg("/handler/", "/handler/some/arg") => "some/arg"
func ParseURLArg(handler, url string) (string, error) {
	l := len(handler)
	if len(url) < len(handler) {
		return "", errors.New("invalid input")
	}

	return url[l:], nil
}
