package main

import (
	"flag"
	"fmt"
)

func scan(path string) {
	fmt.Printf("scan: %s", path)
	repos := recursiveScanFolder(path)
	filePath := getDotFilePath()
	addNewSliceElementsToFile(filePath, repos)
	fmt.Printf("Done")

}

func stats(email string) {
	fmt.Printf("stats: %s \n", email)
	commits := processRepositories(email)
	printCommitsStats(commits)
}

func main() {
	// Define cli flags
	var folder string
	var email string
	flag.StringVar(&email, "email", "your@email.com", "the email to scan")
	flag.StringVar(&folder, "add", "", "add a new folder to scan for Git repos")

	flag.Parse()

	if folder != "" {
		scan(folder)
		return
	}

	stats(email)
}
