package msggeneratorimpl

import (
	"context"
	"fmt"
	srv "github.com/Dima191/RUTUBE-Task/internal/service/msg_generator"
)

func (s *service) Generate(ctx context.Context, subscriberFullName, celebrantFullName string) (string, error) {
	msg, err := s.geminiCl.GenerateMessage(ctx, fmt.Sprintf("%s. Celebratnt Name: %s. From: %s", queryCongratulationsGenerate, celebrantFullName, subscriberFullName))
	if err != nil {
		return "", srv.ErrInternal
	}

	return msg, nil
}
