package handlerimpl

import (
	"github.com/Dima191/RUTUBE-Task/internal/handlers"
	"github.com/Dima191/RUTUBE-Task/internal/models"
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

func TestEmployee(t *testing.T) {
	h, mockedEmployeeSrv, ctrl := testHandler(t)
	defer ctrl.Finish()

	cases := []struct {
		name               string
		employeeID         string
		expectedStatusCode int
		prepareFunc        func()
	}{
		{
			name:               "OK",
			employeeID:         "1",
			expectedStatusCode: http.StatusOK,
			prepareFunc: func() {
				mockedEmployeeSrv.EXPECT().EmployeeByID(gomock.Any(), gomock.Any()).Return(models.Employee{}, nil)
			},
		},
	}

	router := chi.NewRouter()
	router.Get(handlers.EmployeeURL, h.Employee())

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			c.prepareFunc()

			employeeURL := strings.ReplaceAll(handlers.EmployeeURL, "{"+handlers.EmployeeIDURLParam+"}", c.employeeID)
			req, err := http.NewRequest(http.MethodGet, employeeURL, nil)
			assert.NoError(t, err)

			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, req)

			assert.Equal(t, c.expectedStatusCode, recorder.Code)
		})
	}
}

func TestEmployeeErr(t *testing.T) {
	h, mockedEmployeeSrv, ctrl := testHandler(t)
	defer ctrl.Finish()

	cases := []struct {
		name               string
		employeeID         string
		expectedStatusCode int
		expectedMsg        string
		prepareFunc        func()
	}{
		{
			name:               "INVALID EMPLOYEE ID",
			employeeID:         "invalid employee id",
			expectedStatusCode: http.StatusBadRequest,
			expectedMsg:        handlers.ErrInvalidEmployeeID.Error(),
			prepareFunc: func() {
			},
		},
		{
			name:               "EMPLOYEE NOT FOUND",
			employeeID:         "1",
			expectedStatusCode: http.StatusNotFound,
			expectedMsg:        handlers.ErrEmployeeNotFound.Error(),
			prepareFunc: func() {
				mockedEmployeeSrv.EXPECT().EmployeeByID(gomock.Any(), gomock.Any()).Return(models.Employee{}, rep.ErrEmployeeNotFound)
			},
		},
		{
			name:               "INTERNAL ERROR",
			employeeID:         "1",
			expectedStatusCode: http.StatusInternalServerError,
			expectedMsg:        handlers.ErrInternal.Error(),
			prepareFunc: func() {
				mockedEmployeeSrv.EXPECT().EmployeeByID(gomock.Any(), gomock.Any()).Return(models.Employee{}, srv.ErrInternal)
			},
		},
	}

	router := chi.NewRouter()
	router.Get(handlers.EmployeeURL, h.Employee())

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			c.prepareFunc()

			employeeURL := strings.ReplaceAll(handlers.EmployeeURL, "{"+handlers.EmployeeIDURLParam+"}", c.employeeID)
			req, err := http.NewRequest(http.MethodGet, employeeURL, nil)
			require.NoError(t, err)

			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, req)

			assert.Equal(t, c.expectedStatusCode, recorder.Code)
			assert.Contains(t, recorder.Body.String(), c.expectedMsg)
		})
	}
}
