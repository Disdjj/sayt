package cmd

import (
	"fmt"
	"github.com/Disdjj/sayt/internal"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		url := args[1]
		fmt.Printf("Adding `%s` from `%s`\n", name, url)
		internal.InstallRemoteRepo(name, url)
	},
}

func init() {
	repoCmd.AddCommand(addCmd)
}
