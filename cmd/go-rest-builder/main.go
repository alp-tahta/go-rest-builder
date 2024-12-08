package main

import (
	"github.com/alp-tahta/go-rest-builder/internal/skeleton"
)

func main() {
	// initial Domain Name. It will be changed when user provides a name
	domainName := "test"
	tree := skeleton.BuildProjectSkeleton(domainName)
	skeleton.CreateFolders(tree, "")
	skeleton.CreateFiles(tree, "")
}
