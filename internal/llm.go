package internal

import (
	"context"
	"errors"
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"
	"io"
)

func Complete(prompt *PromptNode, message string) {
	stream, err := client.CreateChatCompletionStream(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: viper.GetString("llm.model"),
			Messages: []openai.ChatCompletionMessage{
				{Role: "system", Content: prompt.SystemMessage},
				{Role: "user", Content: message},
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

	defer func(stream *openai.ChatCompletionStream) {
		_ = stream.Close()
	}(stream)

	channel := make(chan string, 1)
	defer close(channel)

	go ShowInTerminal(channel)

	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		content := response.Choices[0].Delta.Content
		channel <- content

	}
}
