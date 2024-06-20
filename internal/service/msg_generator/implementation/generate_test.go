package msggeneratorimpl

import (
	"context"
	geminiclient "github.com/Dima191/RUTUBE-Task/internal/clients/gemini"
	srv "github.com/Dima191/RUTUBE-Task/internal/service/msg_generator"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestGenerate(t *testing.T) {
	s, mockedGeminiCl, ctrl := testService(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mockedGeminiCl.EXPECT().GenerateMessage(gomock.Any(), gomock.Any()).Return("", nil)

	_, err := s.Generate(ctx, "", "")
	assert.NoError(t, err)
}

func TestGenerateErr(t *testing.T) {
	s, mockedGeminiCl, ctrl := testService(t)
	defer ctrl.Finish()

	cases := []struct {
		name        string
		expectedErr error
		prepareFunc func()
	}{
		{
			name:        "UNEXPECTED RESPONSE FORMAT",
			expectedErr: srv.ErrInternal,
			prepareFunc: func() {
				mockedGeminiCl.EXPECT().GenerateMessage(gomock.Any(), gomock.Any()).Return("", geminiclient.ErrUnexpectedResponseFormat)
			},
		},
		{
			name:        "GENERATION RESPONSE ERROR",
			expectedErr: srv.ErrInternal,
			prepareFunc: func() {
				mockedGeminiCl.EXPECT().GenerateMessage(gomock.Any(), gomock.Any()).Return("", geminiclient.ErrGenerateMessage)
			},
		},
	}

	ctx := context.Background()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			c.prepareFunc()

			_, err := s.Generate(ctx, "", "")
			assert.ErrorIs(t, err, c.expectedErr)
		})
	}
}
