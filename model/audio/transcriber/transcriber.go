package transcriber

import (
	"context"
	"github.com/Fufuhu/go-audio-study/util/logging"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/transcribe"
	"go.uber.org/zap"
)

type AudioTranscriberConfig struct{}

type AudioTranscriber struct {
	config *AudioTranscriberConfig
	client *transcribe.Client
}

func GetAudioTranscriber(c *AudioTranscriberConfig) (*AudioTranscriber, error) {
	logger := logging.GetLogger()
	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		logger.Warn(err.Error())
		return nil, err
	}

	return &AudioTranscriber{
		config: c,
		client: transcribe.NewFromConfig(cfg)}, nil
}

func GetDefaultAudioTranscriberConfig() *AudioTranscriberConfig {
	return &AudioTranscriberConfig{}
}
