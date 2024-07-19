package cmd

import (
	"github.com/Disdjj/sayt/internal"
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "build repo",
	Run: func(cmd *cobra.Command, args []string) {
		internal.BuildRepoConfig()
	},
}

func init() {
	repoCmd.AddCommand(buildCmd)
}
