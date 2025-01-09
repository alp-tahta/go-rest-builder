package config

import (
	"fmt"
	"net/url"
	"strings"
)

type Values struct {
	RawModulePath  string
	ModulePath     string
	RootFolderName string
	DomainName     string
}

// TODO explain
func ExtractRootFolderName(rawModulePath string) (rootFolderName string, err error) {
	rootFolderName = rawModulePath

	if strings.ContainsRune(rawModulePath, '/') {
		parsedURL, err := url.Parse(rawModulePath)
		if err != nil {
			return "", err
		}
		path := parsedURL.Path
		segments := strings.Split(strings.Trim(path, "/"), "/")
		rootFolderName = segments[len(segments)-1]
	}

	return rootFolderName, nil
}

/*
ExtractModulePath extract module path from raw module path,
which will be used for go module creation.
*/
func ExtractModulePath(rawModulePath string) (modulePath string) {
	modulePath = rawModulePath

	if strings.HasPrefix(rawModulePath, "http://") {
		modulePath = strings.TrimPrefix(rawModulePath, "http://")
	} else if strings.HasPrefix(rawModulePath, "https://") {
		modulePath = strings.TrimPrefix(rawModulePath, "https://")
	}

	return modulePath
}

// ReceiveValues receive input from user and sets them to config values
func ReceiveValues(v *Values) {
	// Receive module path
	fmt.Println("Enter module path:\n",
		"Example: https://github.com/username/project-name for remote \n",
		"Or project-name for local:")
	fmt.Scanln(&v.RawModulePath)

	// Receive domain name
	fmt.Println("Enter a domain name(Ex. user): ")
	fmt.Scanln(&v.DomainName)
}
