package main

import (
	"dircli/fileutils"  // Import fileutils package
	"dircli/inpututils" // Import inpututils package
	"fmt"
)

func main() {
	baseFolder := "./test/"

	// Get folder name
	folderName := inpututils.GetInput("Enter the name of the folder you want to create: ")
	fullFolderPath := fileutils.CreateFolder(baseFolder, folderName)

	// Get file name
	fileName := inpututils.GetInput("Enter the name of the file you want to create: ")
	fullFilePath := fileutils.CreateFile(fullFolderPath, fileName)

	// Get content to write
	inputText := inpututils.GetInput("Enter the content you want to write to the file: ")
	fileutils.WriteToFile(fullFilePath, inputText)

	fmt.Println("Content written to the file successfully")

	// List files in the folder
	fmt.Println("================")
	fileutils.ListDir(baseFolder)

	// Rename the folder
	fmt.Println("================")
	newFolderName := inpututils.GetInput("Enter the new name for the folder: ")
	fileutils.RenameFolder(baseFolder, folderName, newFolderName)

	// List files in the folder
	fmt.Println("================")
	fileutils.ListDir(baseFolder)
}
