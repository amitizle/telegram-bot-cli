package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/amitizle/telegram-bot-cli/internal/bot"
)

var messageCmd = &cobra.Command{
	Use:   "message",
	Short: "Send a message to the specified chat ID",
	Run:   sendMessage,
}

func init() {
	rootCmd.AddCommand(messageCmd)
	messageCmd.Flags().StringP("message", "m", "", "The message content to send")
	messageCmd.Flags().Int64P("chatid", "i", 0, "Chat ID to send the message to")
	messageCmd.Flags().StringP("token", "t", "", "Telegram bot token")
	viper.BindPFlag("chatid", messageCmd.Flags().Lookup("chatid"))
	viper.BindPFlag("message", messageCmd.Flags().Lookup("message"))
	viper.BindPFlag("token", messageCmd.Flags().Lookup("token"))
}

func sendMessage(cmd *cobra.Command, args []string) {
	token := viper.GetString("token")
	chatID := viper.GetInt64("chatid")
	message := viper.GetString("message")
	bot, err := bot.New(token)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	bot.Message(chatID, message)
}
