package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display the list of all of your bookmarks",
	Long: `Display the list of all of your bookmarks`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Here is the list of all of your bookmarks:\n")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
