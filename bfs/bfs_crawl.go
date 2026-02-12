package bfs

import (
	"chrismangan/search-crawler/html_parser"
	"fmt"
	"net/http"
)

func BFS(startUrl string, maxPageVisits int) ([]string) {
	visits := 0
	foundLinks := []string{startUrl}
	queue := []string{startUrl}

	for {
		if len(queue) == 0 {
			fmt.Println("Queue exhausted, terminating search")
			break
		}

		if visits >= maxPageVisits {
			fmt.Println("Max page visits reached, terminating search")
			break
		}

		currentUrl := queue[0]
		queue = queue[1:]
		response, err := http.Get(currentUrl)

		if err != nil {
			fmt.Printf("Error HTTP GET: %v\n", err)
			continue
		}

		extractedLinks, err := html_parser.ParseHtml(response.Body)

		if err != nil {
			fmt.Printf("Error parsing links on %s: %v\n", currentUrl, err)
			continue
		}

		foundLinks = append(foundLinks, extractedLinks...)
		queue = append(queue, foundLinks...)

		visits++
	}

	return foundLinks
}
