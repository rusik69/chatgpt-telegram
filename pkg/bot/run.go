package bot

import (
	"log"

	"github.com/rusik69/chatgpt-tg/pkg/env"
	"github.com/rusik69/chatgpt-tg/pkg/openaiclient"
	"github.com/rusik69/chatgpt-tg/pkg/telegramclient"
)

// Run runs the bot.
func Run() {
	updates := telegramclient.GetUpdates()
	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s\n", update.Message.From.UserName, update.Message.Text)
			if !env.EnvInstance.AllowedUsers[update.Message.From.UserName] {
				telegramclient.Send(update, "You are not allowed to use this bot.")
				continue
			}
			message, err := openaiclient.Generate(update.Message.Text)
			if err != nil {
				log.Println(err)
				telegramclient.Send(update, err.Error())
				continue
			}
			log.Printf("[%s chatgpt] %s\n", update.Message.From.UserName, message)
			telegramclient.Send(update, message)
		}
	}
}
