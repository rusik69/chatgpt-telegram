package huggingface

import (
	"encoding/json"
)

// Bert is a bert model.
func Bert(prompt string) (string, error) {
	url := "https://api-inference.huggingface.co/models/bert-base-uncased"
	body, err := json.Marshal(BertRequest{Inputs: prompt + "[MASK]"})
	if err != nil {
		return "", err
	}
	data, err := Post(url, body)
	if err != nil {
		return "", err
	}
	var bertResponse []bertResponse
	err = json.Unmarshal(data, &bertResponse)
	if err != nil {
		return "", err
	}
	return bertResponse[0].Sequence, err
}
