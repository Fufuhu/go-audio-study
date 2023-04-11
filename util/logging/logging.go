package logging

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var logger *zap.Logger

func GetLogger() *zap.Logger {
	if logger != nil {
		return logger
	}

	tempDir := os.TempDir()
	file, _ := os.CreateTemp(tempDir, "go_audio_study.*.log")
	fmt.Printf("log file path: %s", file.Name())

	sink := zapcore.AddSync(file)
	lockSink := zapcore.Lock(sink)
	enc := zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())

	core := zapcore.NewCore(enc, lockSink, zapcore.InfoLevel)
	logger = zap.New(core)
	return logger
}
