package geminiclientimpl

import (
	"context"
	"github.com/Dima191/RUTUBE-Task/internal/clients/gemini"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
	"log/slog"
)

const (
	generativeModel = "gemini-1.5-flash"
)

type client struct {
	*genai.Client
	logger *slog.Logger
}

func New(ctx context.Context, apiKey string, logger *slog.Logger) (geminiclient.Client, error) {
	cl, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		logger.Error("fail to create Gemini client", slog.String("error", err.Error()))
		return nil, err
	}

	internalCl := &client{cl, logger}

	return internalCl, nil
}
