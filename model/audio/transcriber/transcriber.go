package transcriber

import (
	"context"
	"fmt"
	"github.com/Fufuhu/go-audio-study/util/logging"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/transcribe"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"path/filepath"
)

type AudioTranscriberConfig struct{}

type AudioTranscriber struct {
	config *AudioTranscriberConfig
	client *transcribe.Client
}

type TranscribeJob struct {
	JobName string
}

// StartTranscribeFile ローカルファイルのTranscribe処理をじっこうします
func (t *AudioTranscriber) StartTranscribeFile(filePath string) (*TranscribeJob, error) {
	logger := logging.GetLogger()
	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)

	absoluteFilePath, err := filepath.Abs(filePath)
	if err != nil {
		logger.Warn(err.Error())
		return nil, err
	}

	media := &types.Media{
		MediaFileUri: aws.String(absoluteFilePath),
	}

	uid, _ := uuid.NewRandom()
	jobName := uid.String()

	settings := &types.Settings{
		MaxSpeakerLabels:  aws.Int32(1),
		ShowSpeakerLabels: aws.Bool(true),
	}

	resp, err := t.client.StartTranscriptionJob(context.TODO(),
		&transcribe.StartTranscriptionJobInput{
			LanguageCode:         types.LanguageCodeJaJp,
			TranscriptionJobName: aws.String(jobName),
			Settings:             settings,
			Media:                media,
		})

	if err != nil {
		logger.Warn(err.Error())
		return nil, err
	}

	logger.Info(fmt.Sprintf("Successfully submit job(%s)", jobName))

	job := &TranscribeJob{JobName: aws.ToString(resp.TranscriptionJob.TranscriptionJobName)}

	return job, nil
}

// WaitToTranscribe StartTranscribeFileの後に処理が終わるのを待つための関数
func (t *AudioTranscriber) WaitToTranscribe(job *TranscribeJob) error {
	return nil
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
