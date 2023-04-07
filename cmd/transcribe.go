/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/Fufuhu/go-audio-study/model/audio/transcriber"
	"github.com/Fufuhu/go-audio-study/util/logging"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
)

// transcribeCmd represents the transcribe command
var transcribeCmd = &cobra.Command{
	Use:   "transcribe",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := logging.GetLogger()
		defer func(logger *zap.Logger) {
			_ = logger.Sync()
		}(logger)

		t, err := transcriber.GetAudioTranscriber(transcriber.GetDefaultAudioTranscriberConfig())
		if err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}

		ctx := context.Background()
		transcription, err := t.Transcribe(ctx, filePath)
		if err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}
		fmt.Println(transcription)
	},
}

func init() {
	rootCmd.AddCommand(transcribeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// transcribeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// transcribeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	transcribeCmd.Flags().StringVarP(&filePath, "filePath", "f", "", "file path to transcribe")
	_ = transcribeCmd.MarkFlagRequired("filePath")
}
