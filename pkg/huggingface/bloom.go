package huggingface

import "encoding/json"

// Bloom is a bloom model.
func Bloom(prompt string) (string, error) {
	url := "https://api-inference.huggingface.co/models/bigscience/bloom"
	body, err := json.Marshal(BloomRequest{Inputs: prompt, MaxNewTokens: 250})
	if err != nil {
		return "", err
	}
	data, err := Post(url, body)
	if err != nil {
		return "", err
	}
	var bloomResponse []bloomResponse
	err = json.Unmarshal(data, &bloomResponse)
	if err != nil {
		return "", err
	}
	var generatedTest string
	for _, text := range bloomResponse {
		generatedTest += text.GeneratedText
	}
	return generatedTest, err
}
