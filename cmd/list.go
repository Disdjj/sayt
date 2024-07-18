package cmd

import (
	"fmt"
	"github.com/Disdjj/sayt/internal"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list repo",
	Long:  `list repo and category: 'list' / 'list category'`,
	Run: func(cmd *cobra.Command, args []string) {
		println("repos:")
		repos := internal.ListRepo()
		for _, repo := range repos {
			fmt.Println("- ", repo)
		}

		if len(args) > 0 {
			println("\ncategories:")
			switch args[0] {
			case "category":
				for _, repo := range repos {
					println(repo)
					for _, category := range internal.ListCategories(repo) {
						fmt.Println("- ", category)
					}
				}
			}
		}
	},
}

func init() {
	repoCmd.AddCommand(listCmd)
}
