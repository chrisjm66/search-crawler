package http_operations

import (
	"net/url"
)

func IsUrlValid (input string) (bool) {
	_, err := url.ParseRequestURI(input)
	if err != nil {
		return false
	}

	return true
}
