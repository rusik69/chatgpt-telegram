package huggingface

import "encoding/json"

// OpenLLama is a llama model.
func OpenLLama(prompt string) (string, error) {
	url := "https://api-inference.huggingface.co/models/openlm-research/open_llama_7b"
	body, err := json.Marshal(llamaRequest{Inputs: prompt})
	if err != nil {
		return "", err
	}
	data, err := Post(url, body)
	if err != nil {
		return "", err
	}
	var llamaResponse []llamaResponse
	err = json.Unmarshal(data, &llamaResponse)
	if err != nil {
		return "", err
	}
	var generatedTest string
	for _, text := range llamaResponse {
		generatedTest += text.GeneratedText
	}
	return generatedTest, err
}
