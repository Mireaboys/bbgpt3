package api

import (
	"context"

	gogpt "github.com/sashabaranov/go-gpt3"
)

type GPT struct {
	Client *gogpt.Client
}

func NewGPT(token string) *GPT {
	return &GPT{
		Client: gogpt.NewClient(token),
	}
}

func (c *GPT) Get(prompt string) (string, error) {
	ctx := context.Background()
	req := gogpt.CompletionRequest{
		Model:       gogpt.GPT3TextDavinci003,
		MaxTokens:   2048,
		Prompt:      prompt,
		Stream:      false,
		Echo:        false,
		Temperature: float32(0.9),
	}
	response, err := c.Client.CreateCompletion(ctx, req)
	if err != nil {
		return "", err
	}
	return response.Choices[0].Text, nil
}
