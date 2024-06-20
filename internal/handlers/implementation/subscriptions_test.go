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
	"testing"
)

func TestSubscriptions(t *testing.T) {
	h, mockedEmployeeSrv, ctrl := testHandler(t)
	defer ctrl.Finish()

	cases := []struct {
		name             string
		expectStatusCode int
		prepareFunc      func() *http.Request
	}{
		{
			name:             "OK",
			expectStatusCode: http.StatusOK,
			prepareFunc: func() *http.Request {
				req, err := http.NewRequest(http.MethodGet, handlers.SubscriptionsURL, nil)
				require.NoError(t, err)

				ctx := context.WithValue(req.Context(), middlewares.EmployeeIDKeyContext, "1")
				req = req.WithContext(ctx)

				mockedEmployeeSrv.EXPECT().Subscriptions(gomock.Any(), gomock.Any()).Return(nil, nil)

				return req
			},
		},
	}

	router := chi.NewRouter()
	router.Get(handlers.SubscriptionsURL, h.Subscriptions())

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			req := c.prepareFunc()

			recorder := httptest.NewRecorder()

			router.ServeHTTP(recorder, req)

			assert.Equal(t, c.expectStatusCode, recorder.Code)
		})
	}
}

func TestSubscriptionsErr(t *testing.T) {
	h, mockedEmployeeSrv, ctrl := testHandler(t)
	defer ctrl.Finish()

	cases := []struct {
		name             string
		expectStatusCode int
		expectedMsg      string
		prepareFunc      func() *http.Request
	}{
		{
			name:             "NO ID IN CONTEXT ERROR",
			expectStatusCode: http.StatusUnauthorized,
			prepareFunc: func() *http.Request {
				req, err := http.NewRequest(http.MethodGet, handlers.SubscriptionsURL, nil)
				require.NoError(t, err)

				return req
			},
		},
		{
			name:             "INVALID ID IN CONTEXT ERROR",
			expectStatusCode: http.StatusUnauthorized,
			prepareFunc: func() *http.Request {
				req, err := http.NewRequest(http.MethodGet, handlers.SubscriptionsURL, nil)
				require.NoError(t, err)

				ctx := context.WithValue(req.Context(), middlewares.EmployeeIDKeyContext, "invalid id")
				req = req.WithContext(ctx)

				return req
			},
		},
		{
			name:             "INTERNAL ERROR",
			expectStatusCode: http.StatusInternalServerError,
			expectedMsg:      handlers.ErrInternal.Error(),
			prepareFunc: func() *http.Request {
				req, err := http.NewRequest(http.MethodGet, handlers.SubscriptionsURL, nil)
				require.NoError(t, err)

				ctx := context.WithValue(req.Context(), middlewares.EmployeeIDKeyContext, "1")
				req = req.WithContext(ctx)

				mockedEmployeeSrv.EXPECT().Subscriptions(gomock.Any(), gomock.Any()).Return(nil, srv.ErrInternal)

				return req
			},
		},
	}

	router := chi.NewRouter()
	router.Get(handlers.SubscriptionsURL, h.Subscriptions())

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			req := c.prepareFunc()

			recorder := httptest.NewRecorder()

			router.ServeHTTP(recorder, req)

			assert.Equal(t, c.expectStatusCode, recorder.Code)
			assert.Contains(t, recorder.Body.String(), c.expectedMsg)
		})
	}
}
