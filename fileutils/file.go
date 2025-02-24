package fileutils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

// CreateFile creates a file inside the given folder and returns its path
func CreateFile(folderPath, fileName string) string {
	fullFilePath := filepath.Join(folderPath, fileName)
	file, err := os.Create(fullFilePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return ""
	}
	defer file.Close()
	fmt.Printf("File '%s' created successfully\n", fileName)
	return fullFilePath
}

// WriteToFile writes the given content to a file
func WriteToFile(filePath, content string) {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString(content)
	writer.Flush()
}