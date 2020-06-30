package cmd

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// clearallCmd represents the clearall command
var clearallCmd = &cobra.Command{
	Use:   "clearall",
	Short: "Clear all the bookmarks",
	Long: `This command delete all of your added bookmarks. Be careful while using this command.`,
	Run: func(cmd *cobra.Command, args []string) {
		inputCmd := exec.Command("/bin/sh", "-c", "> ~/.bookmarks/bookmarks.txt")
		inputCmd.Stdout = os.Stdout
		inputCmd.Stderr = os.Stderr
    	inputCmd.Run()
	},
}

func init() {
	rootCmd.AddCommand(clearallCmd)
}
