package bot

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rusik69/chatgpt-tg/pkg/env"
	"github.com/rusik69/chatgpt-tg/pkg/huggingface"
	"github.com/rusik69/chatgpt-tg/pkg/openaiclient"
	"github.com/rusik69/chatgpt-tg/pkg/telegramclient"
	"github.com/sashabaranov/go-openai"
)

func HandleMessage(update tgbotapi.Update, d map[string][]openai.ChatCompletionMessage) {
	username := update.Message.From.UserName
	if update.Message != nil {
		message := update.Message.Text
		if !env.EnvInstance.AllowedUsers[username] {
			telegramclient.Send(update, "You are not allowed to use this bot.")
			return
		}
		switch update.Message.Command() {
		case "img":
			message = GetMessage(&update)
			log.Printf("[%s img] %s\n", username, message)
			if message == "" {
				telegramclient.Send(update, "Please provide a prompt for image generator.")
				return
			}
			imgUrl, err := openaiclient.GenerateImage(message)
			if err != nil {
				log.Println(err)
				telegramclient.Send(update, err.Error())
				return
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
				return
			}
			photoName, err := huggingface.StableDiffusion(message)
			if err != nil {
				log.Println(err)
				telegramclient.Send(update, err.Error())
				return
			}
			err = telegramclient.SendPhoto(update, photoName)
			if err != nil {
				log.Println(err)
			}
			os.Remove(photoName)
		case "bloom":
			message = GetMessage(&update)
			log.Printf("[%s bloom] %s\n", username, message)
			if message == "" {
				telegramclient.Send(update, "Please provide a prompt for bloom.")
				return
			}
			reply, err := huggingface.Bloom(message)
			if err != nil {
				log.Println(err)
				telegramclient.Send(update, err.Error())
				return
			}
			log.Printf("[%s bloom] %s", username, reply)
			err = telegramclient.Send(update, reply)
			if err != nil {
				log.Println(err)
			}
		case "bert":
			message = GetMessage(&update)
			log.Printf("[%s bert] %s\n", username, message)
			if message == "" {
				telegramclient.Send(update, "Please provide a prompt for bert.")
				return
			}
			reply, err := huggingface.Bert(message)
			if err != nil {
				log.Println(err)
				telegramclient.Send(update, err.Error())
				return
			}
			log.Printf("[%s bert] %s", username, reply)
			err = telegramclient.Send(update, reply)
			if err != nil {
				log.Println(err)
			}
		case "clear":
			log.Printf("[%s clear]\n", username)
			d[username] = []openai.ChatCompletionMessage{}
			err := telegramclient.Send(update, "Chat history cleared.")
			if err != nil {
				log.Println(err)
			}
		case "help":
			log.Printf("[%s help]\n", username)
			telegramclient.Send(update, "Available commands:\n/dialogue - chatgpt dialogue mod\n/bert - bert llm\n/bloom - bloom llm\n/img - generate an image\n/sd - generate image using StableDiffusion\n/clear - clear chat history\n/help - show this message")
		case "dialogue":
			message = GetMessage(&update)
			log.Printf("[%s dialogue] %s\n", username, message)
			openaiclient.AppendResponse(message, username, &d)
			response, err := openaiclient.ChatGPT(message, username, d[username])
			if err != nil {
				log.Println(err)
				telegramclient.Send(update, err.Error())
				return
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
				return
			}
			log.Printf("[%s chatgpt] %s", username, response)
			err = telegramclient.Send(update, response)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
