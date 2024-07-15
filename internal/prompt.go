package internal

import (
	"fmt"
	"github.com/spf13/viper"
	"io"
	"os"
)

func GetCategoryPrompt(category string) string {
	// Load category prompt from: ../prompts/{category}/system.md
	path := fmt.Sprintf("%s/%s/system.md", viper.GetString("prompts.path"), category)
	// open path and read content
	open, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer open.Close()
	// read all content
	// 读取文件内容
	content, err := io.ReadAll(open)
	if err != nil {
		panic(err)
	}

	return string(content)
}

func GetFileContent(path string) string {
	// open path and read content
	open, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer open.Close()
	// read all content
	// 读取文件内容
	content, err := io.ReadAll(open)
	if err != nil {
		panic(err)
	}

	return string(content)
}
