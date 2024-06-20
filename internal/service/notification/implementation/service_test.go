package notificationimpl

import (
	msggenerator "github.com/Dima191/RUTUBE-Task/internal/service/msg_generator"
	srv "github.com/Dima191/RUTUBE-Task/internal/service/notification"
	smtpmanager "github.com/Dima191/RUTUBE-Task/pkg/smtp_manager"
	stubwriter "github.com/Dima191/RUTUBE-Task/pkg/stub_writer"
	"go.uber.org/mock/gomock"
	"log/slog"
	"testing"
)

func testService(t *testing.T) (s srv.Service, mockedMSGGen *msggenerator.MockService, mockedSMTPManager *smtpmanager.MockManager, ctrl *gomock.Controller) {
	t.Helper()

	ctrl = gomock.NewController(t)

	mockedMSGGen = msggenerator.NewMockService(ctrl)
	mockedSMTPManager = smtpmanager.NewMockManager(ctrl)

	logger := slog.New(slog.NewTextHandler(&stubwriter.Writer{}, nil))

	return New(mockedMSGGen, mockedSMTPManager, logger), mockedMSGGen, mockedSMTPManager, ctrl
}
