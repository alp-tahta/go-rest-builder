package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/alp-tahta/go-rest-builder/internal/config"
	"github.com/alp-tahta/go-rest-builder/internal/skeleton"
	"github.com/alp-tahta/go-rest-builder/internal/writer"
)

func main() {
	v := config.Values{}
	config.ReceiveValues(&v)

	rootFolderName, err := config.ExtractRootFolderName(v.RawModulePath)
	if err != nil {
		log.Fatalln(err)
	}
	modulePath := config.ExtractModulePath(v.RawModulePath)

	v.RootFolderName = rootFolderName
	v.ModulePath = modulePath

	config.IsRemote(&v)
	config.PickRemoteName(&v)

	// Check if folder exist. If it exist, close program with error.
	_, err = os.Stat(v.RootFolderName)
	if !os.IsNotExist(err) {
		log.Fatalf("Error: %s Folder exist. Closing Program.\n", v.RootFolderName)
	}

	// Create Project skeleton with given project name
	tree := skeleton.BuildTree(v.RootFolderName, v.DomainName)

	// Create Folders
	err = skeleton.CreateFolders(tree, "")
	if err != nil {
		os.RemoveAll(v.DomainName)
		log.Fatalln(err)
	}

	// Create Files
	err = skeleton.CreateFiles(tree, "")
	if err != nil {
		os.RemoveAll(v.DomainName)
		log.Fatalln(err)
	}

	// Change to the new project directory
	path, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	projectPath := filepath.Join(path, v.RootFolderName)
	if err := os.Chdir(projectPath); err != nil {
		log.Fatalf("Error changing to project directory: %v\n", err)
	}

	// Initialize a new go.mod file
	cmd := exec.Command("go", "mod", "init", v.DomainName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Error running 'go mod init': %v\n", err)
	}

	// TODO make file paths
	motherFolder := writer.FindMotherFolderOfAFile("service.go", tree)

	motherFolders := writer.FindMotherFoldersOfAFolder(motherFolder, tree)
	motherFolders = append(motherFolders, motherFolder)
	motherFolders = append(motherFolders[:0], motherFolders[1:]...)
	fullPath := fmt.Sprintf("%s", strings.Join(motherFolders, "/"))
	filePath := fullPath + "/" + "service.go"
	err = writer.WritePackageNameToFile(motherFolder, filePath)
	if err != nil {
		log.Println("error is ", err)
	}
}
