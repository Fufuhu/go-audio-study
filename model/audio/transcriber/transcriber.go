package transcriber

type AudioTranscriberConfig struct{}

type AudioTranscriber struct {
	config *AudioTranscriberConfig
}

func GetAudioTranscriber(config *AudioTranscriberConfig) *AudioTranscriber {
	return &AudioTranscriber{config: config}
}

func GetDefaultAudioTranscriberConfig() *AudioTranscriberConfig {
	return &AudioTranscriberConfig{}
}
