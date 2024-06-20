package handlerimpl

import (
	"github.com/Dima191/RUTUBE-Task/internal/handlers"
	srv "github.com/Dima191/RUTUBE-Task/internal/service/employee"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEmployees(t *testing.T) {
	h, mockedEmployeeSrv, ctrl := testHandler(t)
	defer ctrl.Finish()

	cases := []struct {
		name               string
		expectedStatusCode int
		prepareFunc        func()
	}{
		{
			name:               "OK",
			expectedStatusCode: http.StatusOK,
			prepareFunc: func() {
				mockedEmployeeSrv.EXPECT().Employees(gomock.Any()).Return(nil, nil)
			},
		},
	}

	router := chi.NewRouter()
	router.Get(handlers.EmployeesURL, h.Employees())

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			c.prepareFunc()

			req, err := http.NewRequest(http.MethodGet, handlers.EmployeesURL, nil)
			require.NoError(t, err)

			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, req)

			assert.Equal(t, c.expectedStatusCode, recorder.Code)
		})
	}
}

func TestEmployeesErr(t *testing.T) {
	h, mockedEmployeeSrv, ctrl := testHandler(t)
	defer ctrl.Finish()

	cases := []struct {
		name               string
		expectedStatusCode int
		expectedMsg        string
		prepareFunc        func()
	}{
		{
			name:               "INTERNAL ERROR",
			expectedStatusCode: http.StatusInternalServerError,
			expectedMsg:        handlers.ErrInternal.Error(),
			prepareFunc: func() {
				mockedEmployeeSrv.EXPECT().Employees(gomock.Any()).Return(nil, srv.ErrInternal)
			},
		},
	}

	router := chi.NewRouter()
	router.Get(handlers.EmployeesURL, h.Employees())

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			c.prepareFunc()

			req, err := http.NewRequest(http.MethodGet, handlers.EmployeesURL, nil)
			require.NoError(t, err)

			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, req)

			assert.Equal(t, c.expectedStatusCode, recorder.Code)
			assert.Contains(t, recorder.Body.String(), c.expectedMsg)
		})
	}
}
