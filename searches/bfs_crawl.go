package searches

import (
	"chrismangan/search-crawler/html_parser"
	"chrismangan/search-crawler/http_operations"
	"chrismangan/search-crawler/input"
	"fmt"
	"net/http"
	"time"
)

func BFS(config input.Configuration) SearchReturnData {
	httpClient := &http.Client{}
	visits := 0
	foundLinks := []string{config.StartingLink}
	queue := []string{config.StartingLink}

	for {
		if len(queue) == 0 {
			fmt.Println("Queue exhausted, terminating search")
			break
		}

		if visits >= config.MaxPageVisits {
			fmt.Println("Max page visits reached, terminating search")
			break
		}

		currentUrl := queue[0]
		queue = queue[1:]
		req, err := http_operations.CreateHttpGetRequest(currentUrl)

		if err != nil {
			fmt.Printf("Error HTTP GET: %v\n", err)
			continue
		}

		resp, err := httpClient.Do(req)

		if err != nil {
			fmt.Printf("Error HTTP GET: %v\n", err)
			continue
		}

		defer resp.Body.Close()
		extractedLinks, err := html_parser.ParseHtml(resp.Body)

		if err != nil {
			fmt.Printf("Error parsing links on %s: %v\n", currentUrl, err)
			continue
		}

		foundLinks = append(foundLinks, extractedLinks...)
		queue = append(queue, foundLinks...)

		visits++
		time.Sleep(time.Second * time.Duration(1/config.RateLimit))
	}

	returnData := SearchReturnData{
		Links: foundLinks,
		TotalPagesVisited: visits,
	}

	return returnData
}
