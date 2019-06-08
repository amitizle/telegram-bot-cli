package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/amitizle/telegram-bot-cli/internal/bot"
	"github.com/logrusorgru/aurora"
)

var (
	messageCmd = &cobra.Command{
		Use:   "message",
		Short: "Send a message to the specified chat ID",
	}
	textCommand = &cobra.Command{
		Use:   "text",
		Short: "Send a text message",
		Run:   sendTextMessage,
		Args:  cobra.ExactArgs(1),
	}
	documentCommand = &cobra.Command{
		Use:   "doc",
		Short: "Send a document",
		Run:   sendDocumentMessage,
		Args:  cobra.ExactArgs(1),
	}
	photoCommand = &cobra.Command{
		Use:   "photo",
		Short: "Send a photo",
		Run:   sendPhotoMessage,
		Args:  cobra.ExactArgs(1),
	}
)

func init() {
	rootCmd.AddCommand(messageCmd)
	messageCmd.PersistentFlags().Int64P("chatid", "i", 0, "Chat ID to send the message to")
	messageCmd.PersistentFlags().StringP("token", "t", "", "Telegram bot token")
	messageCmd.MarkPersistentFlagRequired("chatid")
	messageCmd.MarkPersistentFlagRequired("token")
	viper.BindPFlag("chatid", messageCmd.PersistentFlags().Lookup("chatid"))
	viper.BindPFlag("token", messageCmd.PersistentFlags().Lookup("token"))
	messageCmd.AddCommand(textCommand)
	messageCmd.AddCommand(documentCommand)
	messageCmd.AddCommand(photoCommand)
}

func sendTextMessage(cmd *cobra.Command, args []string) {
	token := viper.GetString("token")
	chatID := viper.GetInt64("chatid")
	message := args[0] // cobra.ExactArgs(1) is used in the command spec
	bot, err := bot.New(token)
	if err != nil {
		fmt.Println(aurora.Sprintf(aurora.Red("Error while authenticating against Telegram servers: %v\n"), err))
		os.Exit(1)
	}
	bot.Message(chatID, message)
}

func sendDocumentMessage(cmd *cobra.Command, args []string) {
	token := viper.GetString("token")
	chatID := viper.GetInt64("chatid")
	fileToSendPath := args[0] // cobra.ExactArgs(1) is used in the command spec
	bot, err := bot.New(token)
	if err != nil {
		fmt.Println(aurora.Sprintf(aurora.Red("Error while authenticating against Telegram servers: %v\n"), err))
		os.Exit(1)
	}
	bot.Document(chatID, fileToSendPath)
}

func sendPhotoMessage(cmd *cobra.Command, args []string) {
	token := viper.GetString("token")
	chatID := viper.GetInt64("chatid")
	fileToSendPath := args[0] // cobra.ExactArgs(1) is used in the command spec
	bot, err := bot.New(token)
	if err != nil {
		fmt.Println(aurora.Sprintf(aurora.Red("Error while authenticating against Telegram servers: %v\n"), err))
		os.Exit(1)
	}
	bot.Photo(chatID, fileToSendPath)
}
