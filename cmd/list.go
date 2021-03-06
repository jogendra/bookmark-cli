package cmd

import (
	"os"
	"fmt"
	"bufio"
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
		f, _ := os.Open(homeDir + "/.bookmarks/bookmarks.txt")
		defer f.Close()
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			fmt.Println(">", scanner.Text())
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
