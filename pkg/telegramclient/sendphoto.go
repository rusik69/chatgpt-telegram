package telegramclient

import (
	"io/ioutil"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// SendPhoto sends a photo to the user.
func SendPhoto(update tgbotapi.Update, photoPath string) error {
	file, err := os.Open(photoPath)
	if err != nil {
		return err
	}
	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	requestFileData := tgbotapi.FileBytes{Name: photoPath, Bytes: fileBytes}
	photoConfig := tgbotapi.NewPhoto(update.Message.Chat.ID, requestFileData)
	photoConfig.ReplyToMessageID = update.Message.MessageID
	_, err = Client.Send(photoConfig)
	return err
}
