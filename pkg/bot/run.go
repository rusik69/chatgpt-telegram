package bot

import (
	"log"

	"github.com/rusik69/chatgpt-tg/pkg/env"
	"github.com/rusik69/chatgpt-tg/pkg/openaiclient"
	"github.com/rusik69/chatgpt-tg/pkg/telegramclient"
	"github.com/sashabaranov/go-openai"
)

// Run runs the bot.
func Run() {
	d := map[string][]openai.ChatCompletionMessage{}
	updates := telegramclient.GetUpdates()
	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s\n", update.Message.From.UserName, update.Message.Text)
			if !env.EnvInstance.AllowedUsers[update.Message.From.UserName] {
				telegramclient.Send(update, "You are not allowed to use this bot.")
				continue
			}
			message, err := openaiclient.Generate(update.Message.Text, update.Message.From.UserName, &d)
			if err != nil {
				log.Println(err)
				telegramclient.Send(update, err.Error())
				continue
			}
			log.Printf("[%s chatgpt] %s", update.Message.From.UserName, message)
			telegramclient.Send(update, message)
		}
	}
}
