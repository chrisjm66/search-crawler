package http_operations

import "net/url"

func IsUrlValid(input string) bool {
	_, err := url.ParseRequestURI(input)

	return err == nil
}
