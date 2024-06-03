package env

import (
	"errors"
	"os"
	"strings"
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
	allowedUsersStr := os.Getenv("ALLOWED_USERS")
	if allowedUsersStr == "" {
		return nil, errors.New("ALLOWED_USERS is not set")
	}
	allowedUsers := make(map[string]bool)
	for _, user := range strings.Split(allowedUsersStr, ",") {
		allowedUsers[string(user)] = true
	}
	return &Env{
		OpenAIApiToken:   openApiToken,
		TelegramBotToken: telegramBotToken,
		AllowedUsers:     allowedUsers,
	}, nil
}
