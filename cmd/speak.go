/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	htgotts "github.com/hegedustibor/htgo-tts"
	"github.com/hegedustibor/htgo-tts/handlers"
	"github.com/hegedustibor/htgo-tts/voices"
	"github.com/spf13/cobra"
	"os"
)

// speakCmd represents the speak command
var speakCmd = &cobra.Command{
	Use:   "speak",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		speech := htgotts.Speech{Folder: "audio", Language: voices.Japanese, Handler: &handlers.Native{}}
		_ = speech.Speak(sentence)
		defer func() {
			_ = os.RemoveAll("./audio")
		}()
	},
}

var sentence string

func init() {
	rootCmd.AddCommand(speakCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// speakCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// speakCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	speakCmd.Flags().StringVarP(&sentence, "sentence", "s", "", "sentence to speak")
	_ = speakCmd.MarkFlagRequired("sentence")
}
