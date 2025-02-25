package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Folder creation command
var folderCmd = &cobra.Command{
	Use:   "folder",
	Short: "Create a folder",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			fmt.Println("Please provide a folder name: --name <folder_name>")
			return
		}
		err := os.Mkdir(name, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating folder:", err)
			return
		}
		fmt.Println("Folder created:", name)
	},
}

func init() {
	folderCmd.Flags().StringP("name", "n", "", "Name of the folder to create")
	rootCmd.AddCommand(folderCmd.Root())
}
