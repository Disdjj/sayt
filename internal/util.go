package internal

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
)

const projectPath = ".sayt/"

func GetProjectPath() string {
	// Get Home Path
	dir, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	projectPath := fmt.Sprintf("%s/%s", dir, projectPath)
	return projectPath
}
