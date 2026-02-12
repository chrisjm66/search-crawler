package input

import (
	"chrismangan/search-crawler/http_operations"
	"fmt"
)

func GetUserInput() Configuration {
	inputValid := false
	var config Configuration

	for !inputValid {
		fmt.Println("Select Search Type: Enter 0 for DFS based crawl, 1 for BFS based crawl")
		fmt.Scan(&config.SearchType)

		fmt.Println("Would you like to use a SQL database? Enter 0 (no) or 1 (true)")
		fmt.Scan(&config.UseDatabase)

		fmt.Println("Enter starting link to use (must be HTTPS): ")
		fmt.Scan(&config.StartingLink)

		fmt.Println("Enter max page visits: ")
		fmt.Scan(&config.MaxPageVisits)

		fmt.Println("Enter rate limit:")
		fmt.Scan(&config.RateLimit)

		// Validation
		inputValid = true
		if config.SearchType < 0 || config.SearchType > 1 {
			inputValid = false
			fmt.Println("Invaid search type")
		}

		if config.UseDatabase {
			inputValid = false
			fmt.Println("Invalid database selection")
		}

		if !http_operations.IsUrlValid(config.StartingLink) {
			inputValid = false
			fmt.Println("Invalid URL")
		}
	}

	return config
}

type Configuration struct {
	SearchType    int // 0 = BFS, 1 = DFS
	UseDatabase   bool
	StartingLink  string
	DebugMode     bool
	MaxPageVisits int
	RateLimit     int
}
