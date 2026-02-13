package http_operations

import (
	"fmt"
	"net/http"
	"io"
)

func CreateHttpGetRequest(inputUrl string) (*http.Request, error) {
	req, err := http.NewRequest("GET", inputUrl, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "GoSearchCrawlerResearch/1.0")

	return req, nil
}

func GetPageBody(inputUrl string) (io.ReadCloser, error) {
	req, err := CreateHttpGetRequest(inputUrl)

	if err != nil {
		fmt.Printf("Error creating HTTP request: %v\n", err)
		return nil, err
	}

	resp, err := (&http.Client{}).Do(req)

	if err != nil {
		fmt.Printf("Error HTTP GET: %v\n", err)
		return nil, err
	}

	return resp.Body, nil
}
