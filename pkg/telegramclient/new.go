package telegramclient

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rusik69/chatgpt-tg/pkg/env"
)

// New creates a new client.
func New() (*tgbotapi.BotAPI, error) {
	client, err := tgbotapi.NewBotAPI(env.EnvInstance.TelegramBotToken)
	if err != nil {
		return nil, err
	}
	log.Printf("Authorized on account %s", client.Self.UserName)
	return client, nil
}
