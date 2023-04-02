/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/Fufuhu/go-audio-study/model/audio/recorder"
	"github.com/Fufuhu/go-audio-study/util/logging"
	"github.com/zenwerk/go-wave"
	"go.uber.org/zap"
	"math/rand"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// recordCmd represents the record command
var recordCmd = &cobra.Command{
	Use:   "record",
	Short: "record the audio from microphone and save to the specified file",
	Long:  "record the audio from microphone and save to the specified file",
	Run:   record,
}

var filePath string

func record(cmd *cobra.Command, args []string) {
	logger := logging.GetLogger()
	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)

	recorderConfig := recorder.GetAudioRecorderDefaultConfig()
	r := recorder.GetAudioRecorder(recorderConfig)

	read, closeRecorder, err := r.StartRecording()
	// defer closeRecorder()
	if err != nil {
		logger.Info(err.Error())
		os.Exit(1)
	}

	// setup Wave file writer
	waveFile, err := os.Create(filePath)
	if err != nil {
		logger.Warn(err.Error())
		os.Exit(1)
	}

	param := wave.WriterParam{
		Out:           waveFile,
		Channel:       recorderConfig.InputChannel,
		SampleRate:    recorderConfig.SampleRate,
		BitsPerSample: 16, // if 16, change to WriteSample16()
	}
	waveWriter, err := wave.NewWriter(param)
	if err != nil {
		logger.Warn(err.Error())
		os.Exit(1)
	}

	var command string
	go func() {
		for {
			_, _ = fmt.Scan(&command)
			logger.Info(fmt.Sprintf("command is %s", command))
			if command == "stop" {
				logger.Info("stop command is called start to stop recording")
				fmt.Println("Recording stopped")
				err = waveWriter.Close()
				if err != nil {
					logger.Warn(err.Error())
				}
				closeRecorder()
				os.Exit(0)
			}
		}
	}()

	// recording in progress ticker. From good old DOS days.
	ticker := []string{
		"-",
		"\\",
		"/",
		"|",
	}
	rand.NewSource(time.Now().UnixNano())

	for {
		if err = read(); err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}
		fmt.Printf("\rRecording is live now. Say something to your microphone! [%v]", ticker[rand.Intn(len(ticker)-1)])
		if _, err = waveWriter.WriteSample16(r.FramesPerBuffer); err != nil {
			logger.Warn(err.Error())
		}
	}
}

func init() {
	rootCmd.AddCommand(recordCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// recordCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// recordCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	recordCmd.Flags().StringVarP(&filePath, "filePath", "f", "", "file path to create output")
	recordCmd.MarkFlagRequired("filePath")
}
