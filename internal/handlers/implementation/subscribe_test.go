package handlerimpl

import (
	"context"
	"github.com/Dima191/RUTUBE-Task/internal/handlers"
	"github.com/Dima191/RUTUBE-Task/internal/middlewares"
	rep "github.com/Dima191/RUTUBE-Task/internal/repository"
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

func TestSubscribe(t *testing.T) {
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

				url := strings.ReplaceAll(handlers.EmployeeURL+handlers.SubscribeURL, "{"+handlers.EmployeeIDURLParam+"}", "2")

				req, err := http.NewRequest(http.MethodPost, url, nil)
				require.NoError(t, err)

				ctx := context.WithValue(req.Context(), middlewares.EmployeeIDKeyContext, "1")
				req = req.WithContext(ctx)

				mockedEmployeeSrv.EXPECT().Subscribe(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

				return req
			},
		},
	}

	router := chi.NewRouter()
	router.Post(handlers.EmployeeURL+handlers.SubscribeURL, h.Subscribe())

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			req := c.prepareFunc()

			recorder := httptest.NewRecorder()

			router.ServeHTTP(recorder, req)

			assert.Equal(t, c.expectedStatusCode, recorder.Code)
		})
	}
}

func TestSubscribeErr(t *testing.T) {
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
				url := strings.ReplaceAll(handlers.EmployeeURL+handlers.SubscribeURL, "{"+handlers.EmployeeIDURLParam+"}", "2")

				req, err := http.NewRequest(http.MethodPost, url, nil)
				require.NoError(t, err)

				return req
			},
		},
		{
			name:               "INVALID ID IN CONTEXT ERROR",
			expectedStatusCode: http.StatusUnauthorized,
			prepareFunc: func() *http.Request {

				url := strings.ReplaceAll(handlers.EmployeeURL+handlers.SubscribeURL, "{"+handlers.EmployeeIDURLParam+"}", "2")

				req, err := http.NewRequest(http.MethodPost, url, nil)
				require.NoError(t, err)

				ctx := context.WithValue(req.Context(), middlewares.EmployeeIDKeyContext, "1test")
				req = req.WithContext(ctx)

				return req
			},
		},
		{
			name:               "INVALID EMPLOYEE ID IN URL",
			expectedStatusCode: http.StatusBadRequest,
			expectedMsg:        handlers.ErrInvalidEmployeeID.Error(),
			prepareFunc: func() *http.Request {
				url := strings.ReplaceAll(handlers.EmployeeURL+handlers.SubscribeURL, "{"+handlers.EmployeeIDURLParam+"}", "test")

				req, err := http.NewRequest(http.MethodPost, url, nil)
				require.NoError(t, err)

				ctx := context.WithValue(req.Context(), middlewares.EmployeeIDKeyContext, "1")
				req = req.WithContext(ctx)

				return req
			},
		},
		{
			name:               "SELF SUBSCRIPTION ERROR",
			expectedStatusCode: http.StatusBadRequest,
			expectedMsg:        srv.ErrSelfSubscription.Error(),
			prepareFunc: func() *http.Request {

				url := strings.ReplaceAll(handlers.EmployeeURL+handlers.SubscribeURL, "{"+handlers.EmployeeIDURLParam+"}", "1")

				req, err := http.NewRequest(http.MethodPost, url, nil)
				require.NoError(t, err)

				ctx := context.WithValue(req.Context(), middlewares.EmployeeIDKeyContext, "1")
				req = req.WithContext(ctx)

				mockedEmployeeSrv.EXPECT().Subscribe(gomock.Any(), gomock.Any(), gomock.Any()).Return(srv.ErrSelfSubscription)

				return req
			},
		},
		{
			name:               "ALREADY SUBSCRIBED",
			expectedStatusCode: http.StatusBadRequest,
			expectedMsg:        rep.ErrAlreadySubscribed.Error(),
			prepareFunc: func() *http.Request {

				url := strings.ReplaceAll(handlers.EmployeeURL+handlers.SubscribeURL, "{"+handlers.EmployeeIDURLParam+"}", "2")

				req, err := http.NewRequest(http.MethodPost, url, nil)
				require.NoError(t, err)

				ctx := context.WithValue(req.Context(), middlewares.EmployeeIDKeyContext, "1")
				req = req.WithContext(ctx)

				mockedEmployeeSrv.EXPECT().Subscribe(gomock.Any(), gomock.Any(), gomock.Any()).Return(rep.ErrAlreadySubscribed)

				return req
			},
		},
		{
			name:               "INTERNAL ERROR",
			expectedStatusCode: http.StatusInternalServerError,
			prepareFunc: func() *http.Request {

				url := strings.ReplaceAll(handlers.EmployeeURL+handlers.SubscribeURL, "{"+handlers.EmployeeIDURLParam+"}", "2")

				req, err := http.NewRequest(http.MethodPost, url, nil)
				require.NoError(t, err)

				ctx := context.WithValue(req.Context(), middlewares.EmployeeIDKeyContext, "1")
				req = req.WithContext(ctx)

				mockedEmployeeSrv.EXPECT().Subscribe(gomock.Any(), gomock.Any(), gomock.Any()).Return(srv.ErrInternal)

				return req
			},
		},
	}

	router := chi.NewRouter()
	router.Post(handlers.EmployeeURL+handlers.SubscribeURL, h.Subscribe())

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			req := c.prepareFunc()

			recorder := httptest.NewRecorder()

			router.ServeHTTP(recorder, req)

			assert.Equal(t, c.expectedStatusCode, recorder.Code)
		})
	}
}
