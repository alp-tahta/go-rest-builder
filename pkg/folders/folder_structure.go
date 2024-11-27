package folders

import (
	"fmt"
	"os"
	"path/filepath"
)

type Node struct {
	FolderName string
	Children   []*Node
}

// ReturnMainNodeOfTree returns main Node for directory structure
func ReturnMainNodeOfTree() Node {
	pName := Node{
		FolderName: "rest-api",
		Children:   nil,
	}

	cmd := Node{
		FolderName: "cmd",
		Children:   []*Node{&pName},
	}

	handler := Node{
		FolderName: "handler",
		Children:   []*Node{},
	}

	service := Node{
		FolderName: "service",
		Children:   []*Node{},
	}

	repository := Node{
		FolderName: "repository",
		Children:   []*Node{},
	}

	internal := Node{
		FolderName: "internal",
		Children:   []*Node{&handler, &service, &repository},
	}

	mainFolder := Node{
		FolderName: "rest-api",
		Children:   []*Node{&cmd, &internal},
	}

	return mainFolder
}

// CreateFolders creates the folder structure recursively
func CreateFolders(node Node, parentPath string) {
	// Construct the full path for the current folder
	currentPath := filepath.Join(parentPath, node.FolderName)

	// Create the directory
	err := os.MkdirAll(currentPath, 0755) // MkdirAll ensures parent directories are created
	if err != nil {
		fmt.Printf("Error creating folder %s: %v\n", currentPath, err)
		return
	}

	// Recurse into children
	for _, child := range node.Children {
		CreateFolders(*child, currentPath)
	}
}
