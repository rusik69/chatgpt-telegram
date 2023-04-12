package huggingface

import "io/ioutil"

// StableDiffusion is a stable diffusion model.
func StableDiffusion(prompt string) (string, error) {
	url := "https://api-inference.huggingface.co/models/stabilityai/stable-diffusion-2-1"
	data, err := Post(url, prompt)
	// photoTempFile is a file with the image.
	photoTempFile, err := ioutil.TempFile("", "stablediffusion.*.png")
	if err != nil {
		return "", err
	}
	// Write the image to the file.
	_, err = photoTempFile.Write(data)
	if err != nil {
		return "", err
	}
	return photoTempFile.Name(), err
}
