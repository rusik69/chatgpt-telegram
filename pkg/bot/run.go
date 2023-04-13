package bot

import (
	"github.com/rusik69/chatgpt-tg/pkg/telegramclient"
	"github.com/sashabaranov/go-openai"
)

// Run runs the bot.
func Run() {
	d := map[string][]openai.ChatCompletionMessage{}
	updates := telegramclient.GetUpdates()
	for update := range updates {
		if update.Message == nil {
			continue
		}
		go HandleMessage(update, d)
	}
}
