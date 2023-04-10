package bot

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// GetMessage is a function that gets a message from a user.
func GetMessage(update *tgbotapi.Update) string {
	prompt := strings.Join(strings.Split(update.Message.Text, " ")[1:], " ")
	return prompt
}
