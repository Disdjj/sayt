package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// repoCmd represents the repo command
var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "3rd Repo management",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("repo called")
	},
}

func init() {
	rootCmd.AddCommand(repoCmd)
}
