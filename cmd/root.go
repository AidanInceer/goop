package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goop",
	Short: "Goop is a CLI tool to manage files and folders",
}

func Execute() {
	// Execute CLI
	rootCmd.Execute()
}
