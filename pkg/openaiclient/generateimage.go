package openaiclient

import (
	"context"
	"log"
	"time"

	openai "github.com/sashabaranov/go-openai"
)

// GenerateImage generates an image.
func GenerateImage(prompt string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	reqUrl := openai.ImageRequest{
		Prompt:         prompt,
		Size:           openai.CreateImageSize256x256,
		ResponseFormat: openai.CreateImageResponseFormatURL,
		N:              1,
	}
	respUrl, err := Client.CreateImage(ctx, reqUrl)
	if err != nil {
		return "", err
	}
	log.Println(respUrl.Data[0].URL)
	return respUrl.Data[0].URL, nil
}
