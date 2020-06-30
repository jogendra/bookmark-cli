package cmd

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new bookmark to bookmarks list",
	Long: `When you run this command with a link or text, it add that link/text to your bookmarks list`,
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, _ := os.UserHomeDir()
		if _, err := os.Stat(homeDir + "/.bookmarks/bookmarks.txt"); os.IsNotExist(err) {
			os.Mkdir(homeDir + "/.bookmarks", 0700)
			os.Create(homeDir + "/.bookmarks/bookmarks.txt")
		}
		commandString := "echo " + "\"" + args[0] + "\"" + " >> ~/.bookmarks/bookmarks.txt"
		exec.Command("/bin/sh", "-c", commandString).Output()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
