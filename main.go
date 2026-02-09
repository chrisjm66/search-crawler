package main

import (
	"fmt"
	"net/url"
)

func main() {
	var choice int
	var useDatabase bool
	var startingLink string
	inputValid := false

	fmt.Println("Welcome to the Search Crawler!")

	for !inputValid {
		fmt.Println("Select Search Type: Enter 0 for DFS based crawl, 1 for BFS based crawl")
		fmt.Scan(&choice)

		fmt.Println("Would you like to use a SQL database? Enter 0 (no) or 1 (true)")
		fmt.Scan(&useDatabase)

		fmt.Println("Enter starting link to use (must be HTTPS): ")
		fmt.Scan(&startingLink)

		// Validation
		inputValid = true
		if choice < 0 || choice > 1 {
			inputValid = false
		}

		_, err := url.ParseRequestURI(startingLink)
		if err != nil {
			inputValid = false
			fmt.Println("Invalid URL")
		}
	}

	// Setup config for app
	var config Configuration
	config.searchType = choice
	config.useDatabase = false
	config.startingLink = startingLink

	fmt.Print(config)
}

type Configuration struct {
	searchType   int // 0 = BFS, 1 = DFS
	useDatabase  bool
	startingLink string
}
