package main

import (
	"fmt"
	"log"

	"github.com/kojikeneda/statuschecker"
)

func main() {
	// Create a slice of StatusCheckers.
	checkers := []statuschecker.StatusChecker{
		&statuschecker.AtlassianChecker{},
		&statuschecker.StatusIOChecker{},
		// You can add more StatusCheckers here for other services.
	}

	TestURL("https://status.atlassian.com", checkers)
}

func TestURL(url string, checkers []statuschecker.StatusChecker) {
	for _, checker := range checkers {
		if checker.CanHandle(url) {
			status, err := checker.CheckStatus()
			if err != nil {
				log.Println("Error checking status:", err)
				continue
			}
			fmt.Println(status)
			return
		}

	}
	fmt.Println("No checker can handle this URL")
}
