package http

import (
	"net/url"
)

func IsUrlValid (url string) (bool) {
	_, err := url.ParseRequestURI(url)
	if err != nil {
		return false
	}

	return true
}
