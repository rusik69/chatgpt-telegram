package huggingface

// StableDiffusion is a stable diffusion model.
func StableDiffusion(prompt string) (string, error) {
	url := "https://api-inference.huggingface.co/models/stabilityai/stable-diffusion-2-1"
	data, err := Post(url, prompt)
	return data, err
}
