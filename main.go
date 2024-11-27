package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// DEFAULT_PROJECT_NAME is the name for newly created project, If user does not provide a project name value
const DEFAULT_PROJECT_NAME = "web-api"

func main() {
	// Create a directory for the new project
	if err := os.Mkdir(DEFAULT_PROJECT_NAME, 0755); err != nil {
		log.Fatalf("Error creating project directory: %v\n", err)
	}

	// Change to the new project directory
	projectPath := filepath.Join(".", DEFAULT_PROJECT_NAME)
	if err := os.Chdir(projectPath); err != nil {
		log.Fatalf("Error changing to project directory: %v\n", err)
	}

	// Initialize a new go.mod file
	cmd := exec.Command("go", "mod", "init", DEFAULT_PROJECT_NAME)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("Error running 'go mod init': %v\n", err)
	}

	log.Println("Go project created successfully!")
}
