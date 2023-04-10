package bot

import (
	"log"

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
				telegramclient.Send(update, imgUrl)
			case "sd":
				message = GetMessage(&update)
				log.Printf("[%s stablediffusion] %s\n", username, message)
				prompt := GetMessage(&update)
				if prompt == "" {
					telegramclient.Send(update, "Please provide a prompt for stablediffusion.")
					continue
				}
				url, err := huggingface.StableDiffusion(prompt)
				if err != nil {
					log.Println(err)
					telegramclient.Send(update, err.Error())
					continue
				}
				telegramclient.Send(update, url)
			case "clear":
				log.Printf("[%s clear]\n", username)
				d[username] = []openai.ChatCompletionMessage{}
				telegramclient.Send(update, "Chat history cleared.")
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
				telegramclient.Send(update, response)
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
				telegramclient.Send(update, response)
			}
		}
	}
}
