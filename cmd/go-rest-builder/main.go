package main

import "github.com/alp-tahta/go-rest-builder/pkg/folders"

func main() {
	folders.CreateFolders(folders.BuildProjectSkeleton(), "")
}
