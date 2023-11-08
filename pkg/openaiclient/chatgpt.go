package openaiclient

import (
	"context"
	"errors"
	"time"

	openai "github.com/sashabaranov/go-openai"
)

// ChatGPT generates a response.
func ChatGPT(prompt, user string, m []openai.ChatCompletionMessage) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 180*time.Second)
	defer cancel()
	resp, err := Client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:     openai.GPT4TurboPreview,
		Messages:  m,
		MaxTokens: 1024,
	})
	if ctx.Err() == context.DeadlineExceeded {
		return "", errors.New("timeout")
	}
	if err != nil {
		return "", err
	}
	if len(resp.Choices) == 0 {
		return "", errors.New("no response from OpenAI")
	}
	content := resp.Choices[0].Message.Content
	if err != nil {
		return "", err
	}
	return content, nil
}
