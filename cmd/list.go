package cmd

import (
	"github.com/Disdjj/sayt/internal"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list repo",
	Run: func(cmd *cobra.Command, args []string) {
		internal.ListRepo()
	},
}

func init() {
	repoCmd.AddCommand(listCmd)
}
