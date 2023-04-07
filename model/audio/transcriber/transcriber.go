package transcriber

import (
	"context"
	"github.com/Fufuhu/go-audio-study/config"
	"github.com/Fufuhu/go-audio-study/util/logging"
	"github.com/sashabaranov/go-openai"
	"go.uber.org/zap"
	"path/filepath"
)

type AudioTranscriberConfig struct {
	OpenAPIKey string
}

type AudioTranscriber struct {
	config *AudioTranscriberConfig
}

// Transcribe ローカルファイルのTranscribe処理を実行します
func (t *AudioTranscriber) Transcribe(ctx context.Context, filePath string) (string, error) {
	logger := logging.GetLogger()
	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)

	absolutePath, err := filepath.Abs(filePath)
	if err != nil {
		logger.Warn(err.Error())
		return "", err
	}

	client := openai.NewClient(t.config.OpenAPIKey)
	req := openai.AudioRequest{
		Model:    openai.Whisper1,
		FilePath: absolutePath,
	}

	resp, err := client.CreateTranscription(ctx, req)
	if err != nil {
		logger.Warn(err.Error())
		return "", err
	}

	return resp.Text, nil
}

func GetAudioTranscriber(c *AudioTranscriberConfig) (*AudioTranscriber, error) {
	logger := logging.GetLogger()
	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)

	return &AudioTranscriber{config: c}, nil
}

func GetDefaultAudioTranscriberConfig() *AudioTranscriberConfig {
	apiKey := config.GetConfig().OpenAPIKey
	return &AudioTranscriberConfig{
		OpenAPIKey: apiKey,
	}
}
