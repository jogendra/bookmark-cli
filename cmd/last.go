package cmd

import (
	"fmt"
	"os"
	"bufio"
	"github.com/spf13/cobra"
)

// lastCmd represents the last command
var lastCmd = &cobra.Command{
	Use:   "last",
	Short: "Show the last most bookmark added",
	Long: `Show the last most bookmark added from the bookmarks list`,
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, _ := os.UserHomeDir()
		if _, err := os.Stat(homeDir + "/.bookmarks/bookmarks.txt");
		os.IsNotExist(err) {
			fmt.Println("Your bookmarks list is empty")
			return
		}
		f, _ := os.Open(homeDir + "/.bookmarks/bookmarks.txt")
		defer f.Close()
		scanner := bufio.NewScanner(f)
		var lastLine string
		for scanner.Scan() {
			lastLine = scanner.Text()
		}
		fmt.Println(">", lastLine)
	},
}

func init() {
	listCmd.AddCommand(lastCmd)
}
