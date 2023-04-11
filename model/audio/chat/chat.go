package chat

import (
	"context"
	"errors"
	"github.com/Fufuhu/go-audio-study/util/logging"
	"github.com/sashabaranov/go-openai"
	"go.uber.org/zap"
)

type AudioChatConfig struct {
	OpenAPIKey string
}

type AudioChatClient struct {
	config   *AudioChatConfig
	client   *openai.Client
	messages []openai.ChatCompletionMessage
}

func (c *AudioChatClient) Chat(message string) (string, error) {
	logger := logging.GetLogger()
	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)

	c.messages = append(c.messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: message,
	})

	var resp openai.ChatCompletionResponse
	resp, err := c.client.CreateChatCompletion(context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: c.messages,
		},
	)
	if err != nil {
		logger.Warn(err.Error())
		return "", err
	}

	if len(resp.Choices) == 0 {
		err = errors.New("resp.Choices has no element")
		logger.Warn(err.Error())
		return "", err
	}

	content := resp.Choices[0].Message.Content
	c.messages = append(c.messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: message,
	})
	return content, nil
}

func GetAudioChatClient(config *AudioChatConfig) *AudioChatClient {
	client := openai.NewClient(config.OpenAPIKey)
	return &AudioChatClient{
		config:   config,
		client:   client,
		messages: make([]openai.ChatCompletionMessage, 0),
	}
}
