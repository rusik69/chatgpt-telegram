package main

import (
	"github.com/rusik69/chatgpt-tg/pkg/bot"
	"github.com/rusik69/chatgpt-tg/pkg/env"
	"github.com/rusik69/chatgpt-tg/pkg/openaiclient"
	"github.com/rusik69/chatgpt-tg/pkg/telegramclient"
)

func main() {
	envInstance, err := env.Parse()
	if err != nil {
		panic(err)
	}
	env.EnvInstance = envInstance
	openaiClient := openaiclient.New()
	openaiclient.Client = openaiClient
	telegramClient, err := telegramclient.New()
	if err != nil {
		panic(err)
	}
	telegramclient.Client = telegramClient
	bot.Run()
}
