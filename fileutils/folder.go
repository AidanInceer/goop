package fileutils

import (
	"fmt"
	"os"
	"path/filepath"
)

// CreateFolder creates a new folder and returns its path
func CreateFolder(basePath, folderName string) string {
	fullFolderPath := filepath.Join(basePath, folderName)
	os.MkdirAll(fullFolderPath, os.ModePerm)
	fmt.Printf("Folder '%s' created successfully\n", folderName)
	return fullFolderPath
}

func ListDir(path string) {
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
