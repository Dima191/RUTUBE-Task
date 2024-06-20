package geminiclientimpl

import (
	"context"
	geminiclient "github.com/Dima191/RUTUBE-Task/internal/clients/gemini"
	"github.com/google/generative-ai-go/genai"
	"log/slog"
)

func (c *client) GenerateMessage(ctx context.Context, query string) (string, error) {
	model := c.GenerativeModel(generativeModel)
	resp, err := model.GenerateContent(ctx, genai.Text(query))
	if err != nil {
		c.logger.Error("failed to generate message", slog.String("error", err.Error()))
		return "", geminiclient.ErrGenerateMessage
	}

	for _, part := range resp.Candidates[0].Content.Parts {
		if text, ok := part.(genai.Text); ok {
			return string(text), nil
		}
	}

	c.logger.Error("failed to generate message", slog.String("error", "unexpected response format"))
	return "", geminiclient.ErrUnexpectedResponseFormat
}
