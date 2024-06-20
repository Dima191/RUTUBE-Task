package msggeneratorimpl

import (
	geminiclient "github.com/Dima191/RUTUBE-Task/internal/clients/gemini"
	msggenerator "github.com/Dima191/RUTUBE-Task/internal/service/msg_generator"
	stubwriter "github.com/Dima191/RUTUBE-Task/pkg/stub_writer"
	"go.uber.org/mock/gomock"
	"log/slog"
	"testing"
)

func testService(t *testing.T) (s msggenerator.Service, mockedGeminiCl *geminiclient.MockClient, ctrl *gomock.Controller) {
	t.Helper()

	ctrl = gomock.NewController(t)

	mockedGeminiCl = geminiclient.NewMockClient(ctrl)

	logger := slog.New(slog.NewTextHandler(&stubwriter.Writer{}, &slog.HandlerOptions{}))

	return New(mockedGeminiCl, logger), mockedGeminiCl, ctrl
}
