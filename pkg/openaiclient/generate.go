package openaiclient

import (
	"context"

	openai "github.com/sashabaranov/go-openai"
)

// Generate generates a response.
func Generate(prompt, user string, d *map[string][]openai.ChatCompletionMessage) (string, error) {
	(*d)[user] = append((*d)[user], openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: prompt,
	})
	resp, err := Client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model:    openai.GPT3Dot5Turbo,
		Messages: (*d)[user],
	})
	content := resp.Choices[0].Message.Content
	if err != nil {
		return "", err
	}
	(*d)[user] = append((*d)[user], openai.ChatCompletionMessage{
		Role: openai.ChatMessageRoleAssistant, Content: content,
	})
	return content, nil
}
