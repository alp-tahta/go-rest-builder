package skeleton

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Node struct {
	FolderName      string
	ChildrenFolders []*Node
	Files           string
}

// BuildProjectSkeleton returns main Node for project structure
func BuildProjectSkeleton(domainName string) Node {
	pName := Node{
		FolderName:      domainName,
		ChildrenFolders: nil,
		Files:           "main.go",
	}

	cmd := Node{
		FolderName:      "cmd",
		ChildrenFolders: []*Node{&pName},
		Files:           "",
	}

	handler := Node{
		FolderName:      "handler",
		ChildrenFolders: []*Node{},
		Files:           "handler.go",
	}

	service := Node{
		FolderName:      "service",
		ChildrenFolders: []*Node{},
		Files:           "service.go",
	}

	repository := Node{
		FolderName:      "repository",
		ChildrenFolders: []*Node{},
		Files:           "repository.go",
	}

	internal := Node{
		FolderName:      "internal",
		ChildrenFolders: []*Node{&handler, &service, &repository},
		Files:           "",
	}

	mainFolder := Node{
		FolderName:      domainName,
		ChildrenFolders: []*Node{&cmd, &internal},
		Files:           "",
	}

	return mainFolder
}

// CreateFolders creates the folder structure recursively
func CreateFolders(node Node, parentPath string) error {
	// Construct the full path for the current folder
	currentPath := filepath.Join(parentPath, node.FolderName)

	// Create the directory
	err := os.MkdirAll(currentPath, 0755) // MkdirAll ensures parent directories are created
	if err != nil {
		log.Println(err)
		return err
	}

	// Recurse into childrenFolders
	for _, child := range node.ChildrenFolders {
		CreateFolders(*child, currentPath)
	}

	return nil
}

// CreateFiles creates files in directories
func CreateFiles(node Node, parentPath string) error {
	// Construct the full path for the current folder
	currentPath := filepath.Join(parentPath, node.FolderName)

	// Create the file
	if node.Files != "" {
		_, err := os.Create(fmt.Sprintf("%s/%s", currentPath, node.Files))
		if err != nil {
			log.Println(err)
			return err
		}
	}

	// Recurse into childrenFolders
	for _, child := range node.ChildrenFolders {
		CreateFiles(*child, currentPath)
	}

	return nil
}
