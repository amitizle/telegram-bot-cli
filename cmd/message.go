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
	_ = messageCmd.MarkPersistentFlagRequired("chatid")
	_ = messageCmd.MarkPersistentFlagRequired("token")
	_ = viper.BindPFlag("chatid", messageCmd.PersistentFlags().Lookup("chatid"))
	_ = viper.BindPFlag("token", messageCmd.PersistentFlags().Lookup("token"))

	documentCommand.Flags().StringP("caption", "a", "", "Document caption")
	_ = viper.BindPFlag("document_caption", documentCommand.Flags().Lookup("caption"))

	photoCommand.Flags().StringP("caption", "a", "", "Photo caption")
	_ = viper.BindPFlag("photo_caption", photoCommand.Flags().Lookup("caption"))

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
	if err := bot.Message(chatID, message); err != nil {
		fmt.Println(aurora.Sprintf(aurora.Red("Failed sending message: %v\n"), err))
		os.Exit(1)
	}
}

func sendDocumentMessage(cmd *cobra.Command, args []string) {
	token := viper.GetString("token")
	chatID := viper.GetInt64("chatid")
	caption := viper.GetString("document_caption")
	fileToSendPath := args[0] // cobra.ExactArgs(1) is used in the command spec
	bot, err := bot.New(token)
	if err != nil {
		fmt.Println(aurora.Sprintf(aurora.Red("Error while authenticating against Telegram servers: %v\n"), err))
		os.Exit(1)
	}
	if err := bot.Document(chatID, fileToSendPath, caption); err != nil {
		fmt.Println(aurora.Sprintf(aurora.Red("Failed sending document: %v\n"), err))
		os.Exit(1)
	}
}

func sendPhotoMessage(cmd *cobra.Command, args []string) {
	token := viper.GetString("token")
	chatID := viper.GetInt64("chatid")
	caption := viper.GetString("photo_caption")
	fileToSendPath := args[0] // cobra.ExactArgs(1) is used in the command spec
	bot, err := bot.New(token)
	if err != nil {
		fmt.Println(aurora.Sprintf(aurora.Red("Error while authenticating against Telegram servers: %v\n"), err))
		os.Exit(1)
	}
	if err := bot.Photo(chatID, fileToSendPath, caption); err != nil {
		fmt.Println(aurora.Sprintf(aurora.Red("Failed sending photo: %v\n"), err))
		os.Exit(1)
	}
}
