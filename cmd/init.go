package cmd

import (
	"fmt"
	"github.com/Disdjj/sayt/internal"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init config file and basic prompts",
	Run:   initImpl,
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func initImpl(cmd *cobra.Command, args []string) {
	initConfig()
	initBasicPrompt()
}

func initConfig() {
	configStr := internal.ConfigToml
	configPath := internal.GetConfigPath()

	// check if file exists
	if _, err := os.Stat(configPath); err == nil {
		fmt.Println("Config file already exists:", configPath)
		return
	}

	// create directory
	// Create all directories in the path if they do not exist
	dirPath := filepath.Dir(configPath)
	if err := os.MkdirAll(dirPath, os.ModeDir); err != nil {
		fmt.Println(err)
		return
	}

	// Create file
	file, err := os.Create(configPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Close file after writing
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	// Write config to file
	_, err = file.WriteString(configStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Success
	fmt.Println("Config file created: ", configPath)
}

func initBasicPrompt() {
}
