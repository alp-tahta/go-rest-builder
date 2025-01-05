package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alp-tahta/go-rest-builder/internal/config"
	"github.com/alp-tahta/go-rest-builder/internal/skeleton"
	"github.com/alp-tahta/go-rest-builder/internal/writer"
)

func main() {
	v := config.Init()
	config.ReceiveValues(v)

	rootFolderName, err := config.ExtractRootFolderName(v.RawModulePath)
	if err != nil {
		log.Fatalln(err)
	}
	modulePath := config.ExtractModulePath(v.RawModulePath)

	v.RootFolderName = rootFolderName
	v.ModulePath = modulePath

	fmt.Println(v)

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

	// TODO make file paths

	err = writer.WriteToFile(fmt.Sprintf("%s/cmd/%s/main.go", v.RootFolderName, v.DomainName), "package main")
	if err != nil {
		log.Fatalln(err)
	}
}
