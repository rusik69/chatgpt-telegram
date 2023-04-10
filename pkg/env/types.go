package env

// Env represents the environment variables.
type Env struct {
	OpenAIApiToken   string
	TelegramBotToken string
	AllowedUsers     map[string]bool
	HuggingFaceToken string
}

// EnvInstance is the global instance of Env.
var EnvInstance *Env
