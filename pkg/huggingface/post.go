package huggingface

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/rusik69/chatgpt-tg/pkg/env"
)

// Request is a request.
func Post(url, prompt string) (string, error) {
	data, err := json.Marshal(Request{Inputs: prompt})
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+env.EnvInstance.HuggingFaceToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bodyBytes), nil
}
