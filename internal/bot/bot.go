package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
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

func (bot *Bot) Document(chatID int64, filePath string) error {
	documentConfig := tgbotapi.NewDocumentUpload(chatID, filePath)
	_, err := bot.bot.Send(documentConfig)
	return err
}

func (bot *Bot) Photo(chatID int64, filePath string) error {
	photoConfig := tgbotapi.NewPhotoUpload(chatID, filePath)
	photoConfig.Caption = "SUN!"
	_, err := bot.bot.Send(photoConfig)
	return err
}
