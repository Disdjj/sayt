package internal

import (
	"context"
	"errors"
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"
	"io"
)

func CompleteWithFile(path string) {
	stream, err := client.CreateChatCompletionStream(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: viper.GetString("llm.model"),
			Messages: []openai.ChatCompletionMessage{
				{Role: "system", Content: GetCategoryPrompt(Arg.Category)},
				{Role: "user", Content: GetFileContent(path)},
			},
			Stream:      true,
			MaxTokens:   1000,
			Temperature: 0.7,
			TopP:        1,
		},
	)
	if err != nil {
		panic(err)
	}
	defer stream.Close()

	channel := make(chan string, 1)
	go ShowInTerminal(channel)

	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			close(channel)
			break
		}
		content := response.Choices[0].Delta.Content
		channel <- content
	}

}
