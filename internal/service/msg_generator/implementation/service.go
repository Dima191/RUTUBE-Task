package msggeneratorimpl

import (
	geminiclient "github.com/Dima191/RUTUBE-Task/internal/clients/gemini"
	msggenerator "github.com/Dima191/RUTUBE-Task/internal/service/msg_generator"
	"log/slog"
)

type service struct {
	geminiCl geminiclient.Client
	logger   *slog.Logger
}

var (
	queryCongratulationsGenerate = "Write an official birthday greeting for a colleague at work."
)

func New(geminiCl geminiclient.Client, logger *slog.Logger) msggenerator.Service {
	s := &service{
		logger:   logger,
		geminiCl: geminiCl,
	}

	return s
}
