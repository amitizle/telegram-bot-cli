package bot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	token string
	bot   *tgbotapi.BotAPI
}

func New(token string) (*Bot, error) {
	tgbot, err := tgbotapi.NewBotAPI(token)
	newBot := &Bot{
		token: token,
		bot:   tgbot,
	}
	if err != nil {
		return newBot, err
	}
	return newBot, nil
}

func (bot *Bot) Message(chatID int64, message string) error {
	msg := tgbotapi.NewMessage(chatID, message)
	_, err := bot.bot.Send(msg)
	return err
}
