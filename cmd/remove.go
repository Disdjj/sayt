package cmd

import (
	"github.com/Disdjj/sayt/internal"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove repo",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		internal.RemoveRepo(args[0])
	},
}

func init() {
	repoCmd.AddCommand(removeCmd)
}
