package recorder

/*
  #include <stdio.h>
  #include <unistd.h>
  #include <termios.h>
  char getch(){
      char ch = 0;
      struct termios old = {0};
      fflush(stdout);
      if( tcgetattr(0, &old) < 0 ) perror("tcsetattr()");
      old.c_lflag &= ~ICANON;
      old.c_lflag &= ~ECHO;
      old.c_cc[VMIN] = 1;
      old.c_cc[VTIME] = 0;
      if( tcsetattr(0, TCSANOW, &old) < 0 ) perror("tcsetattr ICANON");
      if( read(0, &ch,1) < 0 ) perror("read()");
      old.c_lflag |= ICANON;
      old.c_lflag |= ECHO;
      if(tcsetattr(0, TCSADRAIN, &old) < 0) perror("tcsetattr ~ICANON");
      return ch;
  }
*/
import "C"
import (
	"github.com/Fufuhu/go-audio-study/util/logging"
	"github.com/gordonklaus/portaudio"
	"go.uber.org/zap"
)

type AudioRecorder struct {
	config          *AudioRecorderConfig
	stream          *portaudio.Stream
	FramesPerBuffer []int16
}

type AudioRecorderConfig struct {
	InputChannel  int
	OutputChannel int
	SampleRate    int
}

var audioRecorder *AudioRecorder

func (a *AudioRecorder) Initialize() {

}

func (a *AudioRecorder) StartRecording() (func() error, func(), error) {
	logger := logging.GetLogger()
	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)

	// PortAudioの初期化
	logger.Info("portaudio initialize")
	if err := portaudio.Initialize(); err != nil {
		logger.Warn(err.Error())
		return nil, a.DeferStopRecording, err
	}
	logger.Info("portaudio initialize succeeded")

	stream, err := portaudio.OpenDefaultStream(
		a.config.InputChannel,
		a.config.OutputChannel,
		float64(a.config.SampleRate),
		len(a.FramesPerBuffer),
		a.FramesPerBuffer,
	)
	if err != nil {
		logger.Warn(err.Error())
		return nil, a.DeferStopRecording, nil
	}
	a.stream = stream

	err = stream.Start()
	if err != nil {
		logger.Warn(err.Error())
		return nil, a.DeferStopRecording, nil
	}
	return a.stream.Read, a.DeferStopRecording, nil
}

func (a *AudioRecorder) Read() ([]int16, error) {
	logger := logging.GetLogger()
	if err := a.stream.Read(); err != nil {
		logger.Warn(err.Error())
		return a.FramesPerBuffer, err
	}
	return a.FramesPerBuffer, nil
}

func (a *AudioRecorder) DeferStopRecording() {
	logger := logging.GetLogger()
	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)
	if err := portaudio.Terminate(); err != nil {
		logger.Warn(err.Error())
	}
	if err := a.stream.Close(); err != nil {
		logger.Warn(err.Error())
	}
}

func GetAudioRecorderDefaultConfig() *AudioRecorderConfig {
	return &AudioRecorderConfig{
		InputChannel:  1,
		OutputChannel: 0,
		// SampleRate:    48000,
		SampleRate: 16000,
		// SampleRate:    44100,
	}
}

func GetAudioRecorder(config *AudioRecorderConfig) *AudioRecorder {
	if audioRecorder == nil {
		audioRecorder = &AudioRecorder{
			config:          config,
			FramesPerBuffer: make([]int16, 64),
		}
		audioRecorder.Initialize()
	}
	return audioRecorder
}
