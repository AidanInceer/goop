package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// Folder creation command
var cdCmd = &cobra.Command{
	Use:   "cd",
	Short: "changes the current directory",
	Run: func(cmd *cobra.Command, args []string) {
		dir, _ := cmd.Flags().GetString("cd")
		os.Chdir(dir)
	},
}

func init() {
	cdCmd.Flags().StringP("cd", "c", ".", "relative path to target directory")
	rootCmd.AddCommand(cdCmd.Root())
}
