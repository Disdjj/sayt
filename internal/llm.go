package internal

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"
	"sync"
)

type LLMConfig struct {
	Host  string
	Key   string
	Model string
}

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
	viper.SetConfigFile("./config/sayt.toml")
	viper.SetConfigType("toml")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func CompleteWithFile(path string) (string, error) {
	completion, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: viper.GetString("llm.model"),
			Messages: []openai.ChatCompletionMessage{
				{Role: "system", Content: GetCategoryPrompt(Arg.Category)},
				{Role: "user", Content: GetFileContent(path)},
			},
			MaxTokens:   1000,
			Temperature: 0.7,
			TopP:        1,
		},
	)
	if err != nil {
		return "", err
	}
	return completion.Choices[0].Message.Content, nil
}
