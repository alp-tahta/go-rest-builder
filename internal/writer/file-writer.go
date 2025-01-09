package writer

import (
	"os"
	"slices"

	"github.com/alp-tahta/go-rest-builder/internal/skeleton"
)

// TODO
func BuildFilePath(filename string, folderTree skeleton.Node) (motherFolder string) {
	if slices.Contains(folderTree.Files, filename) {
		// If the file is found in the current folder, return its folder name
		return folderTree.FolderName
	}

	// Recursively check in children folders
	for _, child := range folderTree.ChildrenFolders {
		motherFolder = BuildFilePath(filename, child)
		if motherFolder != "" {
			// If the file was found in a child folder, stop further search
			return motherFolder
		}
	}

	// If the file is not found in any folder, return an empty string
	return ""

	/* 	if slices.Contains(folderTree.Files, filename) {
	   		motherFolder = folderTree.FolderName
	   	} else {
	   		for _, v := range folderTree.ChildrenFolders {
	   			if motherFolder == "" {
	   				BuildFilePath(filename, v)
	   			} else {
	   				break
	   			}
	   		}

	   	}
	   	return motherFolder */
	/* ps := findMotherFolder(motherFolder, folderTree)

	fullpath := fmt.Sprintf("%s\n", strings.Join(ps, "/")) */

}

func FindMotherFolders(currentFolderName string, folderTree skeleton.Node) []string {
	// Check each child folder
	for _, child := range folderTree.ChildrenFolders {
		if child.FolderName == currentFolderName {
			// If found, return the path starting with the current folder
			return []string{folderTree.FolderName}
		}

		// Recursively search in the child folder's subtree
		subPath := FindMotherFolders(currentFolderName, child)
		if len(subPath) > 0 {
			// Prepend the current folderTree name to the path
			return append([]string{folderTree.FolderName}, subPath...)
		}
	}

	// If not found, return an empty slice
	return []string{}
}

func WriteToFile(filename string, content string) error {
	// Create or open the file with appropriate permissions
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write content to the file
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}
