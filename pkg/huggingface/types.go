package huggingface

// SDRequest is a stablediffusion request.
type SDRequest struct {
	Inputs string `json:"inputs"`
}

// BertRequest is a bert request.
type BertRequest struct {
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

// bertResponse is a response from the bert model.
type bertResponse struct {
	Sequence string `json:"sequence"`
}
