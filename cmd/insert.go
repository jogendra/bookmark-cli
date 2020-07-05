package cmd

import (
	"os"
	"github.com/spf13/cobra"
)

// insertCmd represents the insert command
var insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "Add a new bookmark to bookmarks list",
	Long: `When you run this command with a link or text, it add that link/text to your bookmarks list`,
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, _ := os.UserHomeDir()
		bookmarkFile := homeDir + "/.bookmarks/bookmarks.txt"
		if _, err := os.Stat(bookmarkFile);
		os.IsNotExist(err) {
			os.Mkdir(homeDir + "/.bookmarks", 0700)
			os.Create(homeDir + "/.bookmarks/bookmarks.txt")
		}
		f, _ := os.OpenFile(bookmarkFile,
							os.O_APPEND|os.O_WRONLY,
							os.ModeAppend) 
		f.WriteString(args[0] + "\n") 
		f.Close()
	},
}

func init() {
	rootCmd.AddCommand(insertCmd)
}
