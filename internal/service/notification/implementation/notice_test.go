package notificationimpl

import (
	"context"
	"errors"
	msggenerator "github.com/Dima191/RUTUBE-Task/internal/service/msg_generator"
	srv "github.com/Dima191/RUTUBE-Task/internal/service/notification"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestNotice(t *testing.T) {
	s, mockedMSGGen, mockedSMTPManager, ctrl := testService(t)
	defer ctrl.Finish()

	ctx := context.Background()
	mockedMSGGen.EXPECT().Generate(gomock.Any(), gomock.Any(), gomock.Any()).Return("", nil)
	mockedSMTPManager.EXPECT().SendMail(gomock.Any(), gomock.Any()).Return(nil)

	err := s.Notice(ctx, "", "", "", "")

	assert.NoError(t, err)
}

func TestNoticeErr(t *testing.T) {
	s, mockedMSGGen, mockedSMTPManager, ctrl := testService(t)
	defer ctrl.Finish()

	cases := []struct {
		name        string
		expectedErr error
		prepareFunc func()
	}{
		{
			name:        "PROCESSING MESSAGE GENERATOR ERROR",
			expectedErr: nil,
			prepareFunc: func() {
				mockedMSGGen.EXPECT().Generate(gomock.Any(), gomock.Any(), gomock.Any()).Return("", msggenerator.ErrInternal)
				mockedSMTPManager.EXPECT().SendMail(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			name:        "PROCESSING MAIL SENDING ERROR",
			expectedErr: srv.ErrInternal,
			prepareFunc: func() {
				mockedMSGGen.EXPECT().Generate(gomock.Any(), gomock.Any(), gomock.Any()).Return("", nil)
				mockedSMTPManager.EXPECT().SendMail(gomock.Any(), gomock.Any()).Return(errors.New("failed to send email"))
			},
		},
	}

	ctx := context.Background()

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			c.prepareFunc()
			err := s.Notice(ctx, "", "", "", "")
			assert.ErrorIs(t, c.expectedErr, err)
		})
	}
}
