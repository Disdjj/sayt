package internal

import (
	"fmt"
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"
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
	viper.SetConfigFile(GetConfigPath())
	viper.SetConfigType("toml")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func Init() {
	InitConfig()
	InitClient()
}
