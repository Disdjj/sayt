package cmd

import (
	"fmt"
	"github.com/Disdjj/sayt/internal"
	"github.com/spf13/cobra"
	"os"
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
	println("len CategoryMap", len(internal.CategoryMap))
	if internal.CategoryMap[promptName] == nil {
		fmt.Printf("Prompt `%s` not found\n", promptName)
		return
	}
	prompt := internal.CategoryMap[promptName]

	switch {
	case len(args) == 2: // user input
		internal.Complete(prompt, args[1])
	default: // Get from stdin
		internal.CompleteStream(prompt, os.Stdin)
	}

}

func init() {
	rootCmd.AddCommand(useCmd)
	useCmd.Flags().StringP("file", "f", "", "file to use")
}
