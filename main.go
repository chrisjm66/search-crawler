package main

import (
	"chrismangan/search-crawler/input"
	"chrismangan/search-crawler/searches"
	"fmt"
	"chrismangan/search-crawler/slice_operations"
	"time"
)

func main() {
	debugMode := true

	fmt.Println("Welcome to the Search Crawler!")

	if debugMode {
		fmt.Println("Debug mode is ACTIVE")
	}

	config := input.GetUserInput()

	if config.DebugMode {
		fmt.Print(config)
	}

	var searchReturnData searches.SearchReturnData
	startNanoTime := time.Now().Nanosecond()
	switch config.SearchType {
	// DFS
	case 0:
		fmt.Println("Not implemented")
	// BFS
	case 1:
		searchReturnData = searches.BFS(config)
	}
	endNanoTime := time.Now().Nanosecond()

	searchReturnData.Links = slice_operations.RemoveDuplicates(searchReturnData.Links)
	fmt.Println("Links found:")
	for _, v := range searchReturnData.Links {
		fmt.Printf("\t%s", v)
	}

	fmt.Printf("Search complete:\n\tTime taken: %d\n", startNanoTime-endNanoTime)
	fmt.Printf("\nTotal pages scraped: %d\nTotal Links Found: %d", searchReturnData.TotalPagesVisited, len(searchReturnData.Links))
}


