package skeleton

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/alp-tahta/go-rest-builder/internal/constants/directory"
	"github.com/alp-tahta/go-rest-builder/internal/constants/file"
)

type Node struct {
	FolderName      string
	ChildrenFolders []Node
	Files           []string
}

// BuildTree returns main Node for project structure
func BuildTree(rootFolderName, domainName string) Node {
	pName := Node{
		FolderName:      domainName,
		ChildrenFolders: nil,
		Files:           []string{file.MAIN},
	}

	cmd := Node{
		FolderName:      directory.CMD,
		ChildrenFolders: []Node{pName},
		Files:           []string{},
	}

	handler := Node{
		FolderName:      directory.HANDLER,
		ChildrenFolders: []Node{},
		Files:           []string{file.HANDLER},
	}

	service := Node{
		FolderName:      directory.SERVICE,
		ChildrenFolders: []Node{},
		Files:           []string{file.SERVICE},
	}

	repository := Node{
		FolderName:      directory.REPOSITORY,
		ChildrenFolders: []Node{},
		Files:           []string{file.REPOSITORY},
	}

	dto := Node{
		FolderName:      directory.DTO,
		ChildrenFolders: []Node{},
		Files:           []string{file.DTO},
	}

	internal := Node{
		FolderName:      directory.INTERNAL,
		ChildrenFolders: []Node{handler, service, repository, dto},
		Files:           []string{},
	}

	mainFolder := Node{
		FolderName:      rootFolderName,
		ChildrenFolders: []Node{cmd, internal},
		Files:           []string{},
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
		CreateFolders(child, currentPath)
	}

	return nil
}

// CreateFiles creates files in directories
func CreateFiles(node Node, parentPath string) error {
	// Construct the full path for the current folder
	currentPath := filepath.Join(parentPath, node.FolderName)

	// Create the file
	if node.Files != nil {
		for _, v := range node.Files {
			_, err := os.Create(fmt.Sprintf("%s/%s", currentPath, v))
			if err != nil {
				log.Println(err)
				return err
			}
		}
	}

	// Recurse into childrenFolders
	for _, child := range node.ChildrenFolders {
		CreateFiles(child, currentPath)
	}

	return nil
}
