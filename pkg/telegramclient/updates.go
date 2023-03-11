package telegramclient

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func GetUpdates() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := Client.GetUpdatesChan(u)
	return updates
}
