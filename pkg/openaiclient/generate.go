package openaiclient

import (
	"context"
	"errors"
	"time"

	openai "github.com/sashabaranov/go-openai"
)

// Generate generates a response.
func Generate(prompt, user string, d *map[string][]openai.ChatCompletionMessage) (string, error) {
	appendResponse(prompt, user, d)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	resp, err := Client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:    openai.GPT3Dot5Turbo,
		Messages: (*d)[user],
	})
	if ctx.Err() == context.DeadlineExceeded {
		return "", errors.New("timeout")
	}
	if len(resp.Choices) == 0 {
		return "", errors.New("no response from OpenAI")
	}
	content := resp.Choices[0].Message.Content
	if err != nil {
		return "", err
	}
	appendResponse(content, user, d)
	return content, nil
}
