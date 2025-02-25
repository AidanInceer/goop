package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// Folder creation command
var treeCmd = &cobra.Command{
	Use:   "tree",
	Short: "displays a tree of the current directory",
	Run: func(cmd *cobra.Command, args []string) {
		grow, _ := cmd.Flags().GetString("grow")
		if grow == "" {
			grow = "."
		}
		displayTree(grow, 0, false)
	},
}

func init() {
	treeCmd.Flags().StringP("grow", "g", ".", "path to the directory to display")
	rootCmd.AddCommand(treeCmd.Root())
}

func displayTree(path string, level int, last bool) {
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for i, file := range files {
		// Ignore hidden files (files starting with a dot)
		if file.Name()[0] == '.' {
			continue
		}
		sSpace := strings.Repeat("│   ", level)

		var prefix string
		if i == len(files)-1 {
			prefix = "└── " // For the last item, we use "└──"
		} else {
			prefix = "├── " // For others, use "├──"
		}

		if file.IsDir() {
			fmt.Println(sSpace + prefix + file.Name())
			displayTree(path+"/"+file.Name(), level+1, i == len(files)-1)
		} else if !last {

			fmt.Println(sSpace + prefix + file.Name())
		} else {
			dBase := strings.Repeat("│   ", level-1)
			fmt.Println(dBase + "    " + prefix + file.Name() + "  ")
		}
	}

}
