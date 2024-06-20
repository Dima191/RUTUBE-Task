package handlerimpl

import (
	"context"
	"github.com/Dima191/RUTUBE-Task/internal/handlers"
	"github.com/Dima191/RUTUBE-Task/internal/middlewares"
	srv "github.com/Dima191/RUTUBE-Task/internal/service/employee"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUnsubscribe(t *testing.T) {
	h, mockedEmployeeSrv, ctrl := testHandler(t)
	defer ctrl.Finish()

	cases := []struct {
		name               string
		expectedStatusCode int
		prepareFunc        func() *http.Request
	}{
		{
			name:               "OK",
			expectedStatusCode: http.StatusNoContent,
			prepareFunc: func() *http.Request {
				url := strings.ReplaceAll(handlers.EmployeeURL+handlers.UnsubscribeURL, "{"+handlers.EmployeeIDURLParam+"}", "2")

				req, err := http.NewRequest(http.MethodPost, url, nil)
				require.NoError(t, err)

				ctx := context.WithValue(req.Context(), middlewares.EmployeeIDKeyContext, "1")
				req = req.WithContext(ctx)

				mockedEmployeeSrv.EXPECT().Unsubscribe(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

				return req
			},
		},
	}

	router := chi.NewRouter()
	router.Post(handlers.EmployeeURL+handlers.UnsubscribeURL, h.Unsubscribe())

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			req := c.prepareFunc()

			recorder := httptest.NewRecorder()

			router.ServeHTTP(recorder, req)

			assert.Equal(t, c.expectedStatusCode, recorder.Code)
		})
	}
}

func TestUnsubscribeErr(t *testing.T) {
	h, mockedEmployeeSrv, ctrl := testHandler(t)
	defer ctrl.Finish()

	cases := []struct {
		name               string
		expectedStatusCode int
		expectedMsg        string
		prepareFunc        func() *http.Request
	}{
		{
			name:               "NO ID IN CONTEXT ERROR",
			expectedStatusCode: http.StatusUnauthorized,
			prepareFunc: func() *http.Request {
				url := strings.ReplaceAll(handlers.EmployeeURL+handlers.UnsubscribeURL, "{"+handlers.EmployeeIDURLParam+"}", "2")

				req, err := http.NewRequest(http.MethodPost, url, nil)
				require.NoError(t, err)

				return req
			},
		},
		{
			name:               "INVALID ID IN CONTEXT ERROR",
			expectedStatusCode: http.StatusUnauthorized,
			prepareFunc: func() *http.Request {
				url := strings.ReplaceAll(handlers.EmployeeURL+handlers.UnsubscribeURL, "{"+handlers.EmployeeIDURLParam+"}", "2")

				req, err := http.NewRequest(http.MethodPost, url, nil)
				require.NoError(t, err)

				ctx := context.WithValue(req.Context(), middlewares.EmployeeIDKeyContext, "invalid id")
				req = req.WithContext(ctx)

				return req
			},
		},
		{
			name:               "INTERNAL ERROR",
			expectedStatusCode: http.StatusInternalServerError,
			expectedMsg:        handlers.ErrInternal.Error(),
			prepareFunc: func() *http.Request {
				url := strings.ReplaceAll(handlers.EmployeeURL+handlers.UnsubscribeURL, "{"+handlers.EmployeeIDURLParam+"}", "2")

				req, err := http.NewRequest(http.MethodPost, url, nil)
				require.NoError(t, err)

				ctx := context.WithValue(req.Context(), middlewares.EmployeeIDKeyContext, "1")
				req = req.WithContext(ctx)

				mockedEmployeeSrv.EXPECT().Unsubscribe(gomock.Any(), gomock.Any(), gomock.Any()).Return(srv.ErrInternal)

				return req
			},
		},
	}

	router := chi.NewRouter()
	router.Post(handlers.EmployeeURL+handlers.UnsubscribeURL, h.Unsubscribe())

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			req := c.prepareFunc()

			recorder := httptest.NewRecorder()

			router.ServeHTTP(recorder, req)

			assert.Equal(t, c.expectedStatusCode, recorder.Code)
			assert.Contains(t, recorder.Body.String(), c.expectedMsg)
		})
	}
}
