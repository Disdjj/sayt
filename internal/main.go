package internal

import (
	"fmt"
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"
	"os"
	"sync"
)

var client *openai.Client
var once sync.Once

func InitClient() {
	config := openai.DefaultConfig(viper.GetString("llm.api_key"))
	config.BaseURL = viper.GetString("llm.api_host")
	once.Do(
		func() {
			c := openai.NewClientWithConfig(config)
			client = c
		},
	)
}

func InitConfig() {
	// 获取~目录位置
	osUserHome, _ := os.UserHomeDir()
	viper.SetConfigFile(osUserHome + "/.sayt/config.toml")
	viper.SetConfigType("toml")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
