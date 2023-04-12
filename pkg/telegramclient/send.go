package telegramclient

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// Send sends a message to the user.
func Send(update tgbotapi.Update, message string) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ReplyToMessageID = update.Message.MessageID
	_, err := Client.Send(msg)
	return err
}
