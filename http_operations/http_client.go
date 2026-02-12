package http_operations

import "net/http"

func CreateHttpGetRequest(inputUrl string) (*http.Request, error) {
	req, err := http.NewRequest("GET", inputUrl, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "GoSearchCrawlerResearch/1.0")

	return req, nil
}
