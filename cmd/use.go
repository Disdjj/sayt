package cmd

import (
	"fmt"
	"github.com/Disdjj/sayt/internal"
	"github.com/spf13/cobra"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "sayt use <prompt name: like assistant> [options]",
	Args:  cobra.MinimumNArgs(1),
	Run:   useImpl,
}

func useImpl(cmd *cobra.Command, args []string) {
	internal.Init()
	promptName := args[0]
	if internal.CategoryMap[promptName] == nil {
		fmt.Printf("Prompt `%s` not found\n", promptName)
		return
	}
	prompt := internal.CategoryMap[promptName]

	if len(args) >= 2 {
		internal.Complete(prompt, args[1])
	}

}

func init() {
	rootCmd.AddCommand(useCmd)
	useCmd.Flags().StringP("file", "f", "", "file to use")
}
