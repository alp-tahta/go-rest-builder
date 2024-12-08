package main

import (
	"log"
	"os"

	"github.com/alp-tahta/go-rest-builder/internal/skeleton"
)

func main() {
	// Initial Domain Name. It will be changed when user provides a name
	domainName := "test"

	// Check if folder exist. If it exist, close program with error.
	_, err := os.Stat(domainName)
	if !os.IsNotExist(err) {
		log.Fatalf("Error: %s Folder exist. Closing Program.\n", domainName)
	}

	// Create Project skeleton with given project name
	tree := skeleton.BuildProjectSkeleton(domainName)

	// Create Folders
	err = skeleton.CreateFolders(tree, "")
	if err != nil {
		os.RemoveAll(domainName)
		log.Fatalln(err)
	}

	// Create Files
	err = skeleton.CreateFiles(tree, "")
	if err != nil {
		os.RemoveAll(domainName)
		log.Fatalln(err)
	}
}
