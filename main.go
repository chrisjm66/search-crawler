package main

import (
	"chrismangan/search-crawler/bfs"
	"chrismangan/search-crawler/http_operations"
	"fmt"
	"time"
)

func main() {
	debugMode := true
	var choice int
	var useDatabase int
	var startingLink string
	var maxPageVisits int
	inputValid := false

	fmt.Println("Welcome to the Search Crawler!")

	if debugMode {
		fmt.Println("Debug mode is ACTIVE")
	}
	for !inputValid {
		fmt.Println("Select Search Type: Enter 0 for DFS based crawl, 1 for BFS based crawl")
		fmt.Scan(&choice)

		fmt.Println("Would you like to use a SQL database? Enter 0 (no) or 1 (true)")
		fmt.Scan(&useDatabase)

		fmt.Println("Enter starting link to use (must be HTTPS): ")
		fmt.Scan(&startingLink)

		fmt.Println("Enter max page visits: ")
		fmt.Scan(&maxPageVisits)

		// Validation
		inputValid = true
		if choice < 0 || choice > 1 {
			inputValid = false
			fmt.Println("Invaid search type")
		}

		if useDatabase < 0 || useDatabase > 1 {
			inputValid = false
			fmt.Println("Invalid database selection")
		}

		if !http_operations.IsUrlValid(startingLink) {
			inputValid = false
			fmt.Println("Invalid URL")
		}
	}

	// Setup config for app
	var config Configuration
	config.searchType = choice
	config.useDatabase = false
	config.startingLink = startingLink
	config.debugMode = debugMode
	config.maxPageVisits = maxPageVisits

	if config.debugMode {
		fmt.Print(config)
	}

	startNanoTime := time.Now().Nanosecond()
	switch(config.searchType) {
		// DFS
		case 0:
			fmt.Println("Not implemented")
			break
		// BFS
		case 1:
			bfs.BFS(config.startingLink, config.maxPageVisits)
			break
	}
	endNanoTime := time.Now().Nanosecond()

	fmt.Printf("Search complete:\n\tTime taken: %d", endNanoTime - startNanoTime)
}

type Configuration struct {
	searchType   int // 0 = BFS, 1 = DFS
	useDatabase  bool
	startingLink string
	debugMode    bool
	maxPageVisits int
}
