package cmd

import (
	"os"
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display the list of all of your bookmarks",
	Long: `Display the list of all of your bookmarks`,
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, _ := os.UserHomeDir()
		if _, err := os.Stat(homeDir + "/.bookmarks/bookmarks.txt"); os.IsNotExist(err) {
			fmt.Println("Your bookmarks list is empty")
			return
		}
		inputCmd := exec.Command("/bin/sh", "-c", "cat ~/.bookmarks/bookmarks.txt")
		inputCmd.Stdout = os.Stdout
		inputCmd.Stderr = os.Stderr
    	inputCmd.Run()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
