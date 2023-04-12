package bot

import (
	"log"
	"os"

	"github.com/rusik69/chatgpt-tg/pkg/env"
	"github.com/rusik69/chatgpt-tg/pkg/huggingface"
	"github.com/rusik69/chatgpt-tg/pkg/openaiclient"
	"github.com/rusik69/chatgpt-tg/pkg/telegramclient"
	"github.com/sashabaranov/go-openai"
)

// Run runs the bot.
func Run() {
	d := map[string][]openai.ChatCompletionMessage{}
	updates := telegramclient.GetUpdates()
	for update := range updates {
		username := update.Message.From.UserName
		if update.Message != nil {
			message := update.Message.Text
			if !env.EnvInstance.AllowedUsers[username] {
				telegramclient.Send(update, "You are not allowed to use this bot.")
				continue
			}
			switch update.Message.Command() {
			case "img":
				message = GetMessage(&update)
				log.Printf("[%s img] %s\n", username, message)
				if message == "" {
					telegramclient.Send(update, "Please provide a prompt for image generator.")
					continue
				}
				imgUrl, err := openaiclient.GenerateImage(message)
				if err != nil {
					log.Println(err)
					telegramclient.Send(update, err.Error())
					continue
				}
				err = telegramclient.Send(update, imgUrl)
				if err != nil {
					log.Println(err)
				}
			case "sd":
				message = GetMessage(&update)
				log.Printf("[%s stablediffusion] %s\n", username, message)
				if message == "" {
					telegramclient.Send(update, "Please provide a prompt for stablediffusion.")
					continue
				}
				photoName, err := huggingface.StableDiffusion(message)
				if err != nil {
					log.Println(err)
					telegramclient.Send(update, err.Error())
					continue
				}
				err = telegramclient.SendPhoto(update, photoName)
				if err != nil {
					log.Println(err)
				}
				os.Remove(photoName)
			case "clear":
				log.Printf("[%s clear]\n", username)
				d[username] = []openai.ChatCompletionMessage{}
				err := telegramclient.Send(update, "Chat history cleared.")
				if err != nil {
					log.Println(err)
				}
			case "help":
				log.Printf("[%s help]\n", username)
				telegramclient.Send(update, "Available commands:\n/dialogue - chatgpt dialogue mod\n/img - generate an image\n/sd - generate image using StableDiffusion\n/clear - clear chat history\n/help - show this message")
			case "dialogue":
				message = GetMessage(&update)
				log.Printf("[%s dialogue] %s\n", username, message)
				openaiclient.AppendResponse(message, username, &d)
				response, err := openaiclient.ChatGPT(message, username, d[username])
				if err != nil {
					log.Println(err)
					telegramclient.Send(update, err.Error())
					continue
				}
				openaiclient.AppendResponse(response, username, &d)
				log.Printf("[%s dialogue] %s", username, response)
				err = telegramclient.Send(update, response)
				if err != nil {
					log.Println(err)
				}
			default:
				log.Printf("[%s] %s\n", username, message)
				if len(d[username]) > 0 {
					d[username] = []openai.ChatCompletionMessage{}
				}
				openaiclient.AppendResponse(message, username, &d)
				response, err := openaiclient.ChatGPT(message, username, d[username])
				if err != nil {
					log.Println(err)
					telegramclient.Send(update, err.Error())
					continue
				}
				log.Printf("[%s chatgpt] %s", username, response)
				err = telegramclient.Send(update, response)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}
}
