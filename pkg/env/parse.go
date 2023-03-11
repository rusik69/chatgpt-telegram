package env

import (
	"errors"
	"os"
)

// Parse parses the environment variables and returns an Env struct.
func Parse() (*Env, error) {
	openApiToken := os.Getenv("OPENAI_API_KEY")
	if openApiToken == "" {
		return nil, errors.New("OPENAI_API_KEY is not set")
	}
	telegramBotToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if telegramBotToken == "" {
		return nil, errors.New("TELEGRAM_BOT_TOKEN is not set")
	}
	return &Env{
		OpenAIApiToken:   openApiToken,
		TelegramBotToken: telegramBotToken,
	}, nil
}
