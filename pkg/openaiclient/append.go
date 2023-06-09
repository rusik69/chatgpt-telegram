package openaiclient

import openai "github.com/sashabaranov/go-openai"

// AppendResponse appends the response to the messages.
func AppendResponse(content, user string, d *map[string][]openai.ChatCompletionMessage) {
	if len((*d)[user]) == 6 {
		newSlice := (*d)[user][1:]
		(*d)[user] = newSlice
	}
	(*d)[user] = append((*d)[user], openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: content,
	})
}
