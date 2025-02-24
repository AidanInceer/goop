package fileutils

import (
	"fmt"
	"os"
	"path/filepath"
)

func RenameFolder(baseFolder string, folder string, newFolder string) {
	var fullPath string = filepath.Join(baseFolder, folder)
	var newName string = filepath.Join(baseFolder, newFolder)

	fmt.Println(fullPath)
	fmt.Println(newName)

	err := os.Rename(fullPath, newName)
	if err != nil {
		fmt.Println("Error renaming folder:", err)
		return
	}
	fmt.Println("Folder renamed successfully")
}
