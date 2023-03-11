package openaiclient

import (
	"github.com/rusik69/chatgpt-tg/pkg/env"
	openai "github.com/sashabaranov/go-openai"
)

// New creates a new client.
func New() *openai.Client {
	client := openai.NewClient(env.EnvInstance.OpenAIApiToken)
	return client
}
