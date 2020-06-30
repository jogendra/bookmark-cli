package cmd

import (
	"os/exec"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a link/text to bookmarks list",
	Long: `When you run this command with a link or text, it add that link/text to your bookmarks list`,
	Run: func(cmd *cobra.Command, args []string) {
		commandString := "echo " + "\"" + args[0] + "\"" + " >> ~/.bookmarks/bookmarks.txt"
		exec.Command("/bin/sh", "-c", commandString).Output()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
