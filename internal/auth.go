package internal

import (
	"errors"
	"net/http"
	"strings"
)

// Extracts an API Key from the headers of an HTTP request
// Example:
// Authorization: ApiKey [api_key]
func getAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("No authentication info found.")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("Malformed authentication header.")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("Malformed first part of authentication header.")
	}

	return vals[1], nil
}
