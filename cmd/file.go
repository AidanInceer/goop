package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// File creation command
var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "Create a file",
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")
		if path == "" {
			fmt.Println("Please provide a file path: --path <file_path>")
			return
		}
		file, err := os.Create(path)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()
		fmt.Println("File created:", path)
	},
}

func init() {
	fileCmd.Flags().StringP("path", "p", "", "Path of the file to create")
	rootCmd.AddCommand(fileCmd)
}
