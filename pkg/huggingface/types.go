package huggingface

// Request is a request.
type SDRequest struct {
	Inputs string `json:"inputs"`
}

// Request is a request.
type BloomRequest struct {
	Inputs       string `json:"inputs"`
	MaxNewTokens int    `json:"max_new_tokens"`
}

// bloomResponse is a response from the bloom model.
type bloomResponse struct {
	GeneratedText string `json:"generated_text"`
}
