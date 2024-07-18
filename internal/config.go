package internal

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
)

const configPath = ".sayt/config.toml"

func GetConfigPath() string {
	// Get Home Path
	dir, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	configPath := fmt.Sprintf("%s/%s", dir, configPath)
	return configPath
}
