package writer

import (
	"fmt"
	"os"
	"slices"

	"github.com/alp-tahta/go-rest-builder/internal/skeleton"
)

// TODO
func FindMotherFolderOfAFile(filename string, folderTree skeleton.Node) (motherFolder string) {
	if slices.Contains(folderTree.Files, filename) {
		// If the file is found in the current folder, return its folder name
		return folderTree.FolderName
	}

	// Recursively check in children folders
	for _, child := range folderTree.ChildrenFolders {
		motherFolder = FindMotherFolderOfAFile(filename, child)
		if motherFolder != "" {
			// If the file was found in a child folder, stop further search
			return motherFolder
		}
	}

	// If the file is not found in any folder, return an empty string
	return ""
}

func FindMotherFoldersOfAFolder(currentFolderName string, folderTree skeleton.Node) []string {
	// Check each child folder
	for _, child := range folderTree.ChildrenFolders {
		if child.FolderName == currentFolderName {
			// If found, return the path starting with the current folder
			return []string{folderTree.FolderName}
		}

		// Recursively search in the child folder's subtree
		subPath := FindMotherFoldersOfAFolder(currentFolderName, child)
		if len(subPath) > 0 {
			// Prepend the current folderTree name to the path
			return append([]string{folderTree.FolderName}, subPath...)
		}
	}

	// If not found, return an empty slice
	return []string{}
}

func buildImport() {

}

func WritePackageNameToFile(pn string, filepath string) error {
	file, err := os.OpenFile(fmt.Sprintf("%s", filepath), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write package name to file
	packageString := fmt.Sprintf("package %s\n\n", pn)
	_, err = file.WriteString(packageString)
	if err != nil {
		return err
	}

	// Write import to file
	importString := fmt.Sprintf("import %s\n", pn)
	_, err = file.WriteString(importString)
	if err != nil {
		return err
	}

	return nil
}
