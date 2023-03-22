package openaiclient

import openai "github.com/sashabaranov/go-openai"

// appendResponse appends the response to the messages.
func appendResponse(content, user string, d *map[string][]openai.ChatCompletionMessage) {
	if len((*d)[user]) == 20 {
		newSlice := (*d)[user][1:]
		(*d)[user] = newSlice
	}
	(*d)[user] = append((*d)[user], openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: content,
	})
}
