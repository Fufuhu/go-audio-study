package config

import (
	"github.com/Fufuhu/go-audio-study/util/logging"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

type ApplicationConfig struct {
	Bucket string `envconfig:"BUCKET"`
}

var config *ApplicationConfig

func GetConfig() *ApplicationConfig {
	logger := logging.GetLogger()
	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)

	if config == nil {
		logger.Info("config is nil. initializing...")
		config = &ApplicationConfig{}
		if err := envconfig.Process("GO_AUDIO_STUDY", config); err != nil {
			logger.Error(err.Error())
			return nil
		}
	}

	return config
}
