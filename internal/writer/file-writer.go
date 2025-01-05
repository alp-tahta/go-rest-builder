package writer

import (
	"os"
)

// TODO
func buildFilePath(filename string) string {
	return ""
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
