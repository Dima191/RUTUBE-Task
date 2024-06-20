package handlerimpl

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/Dima191/RUTUBE-Task/internal/handlers"
	"github.com/Dima191/RUTUBE-Task/internal/models"
	rep "github.com/Dima191/RUTUBE-Task/internal/repository"
	srv "github.com/Dima191/RUTUBE-Task/internal/service/employee"
	"github.com/go-chi/chi"
	"github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignUp(t *testing.T) {
	h, mockedEmployeeSrv, ctrl := testHandler(t)
	defer ctrl.Finish()

	cases := []struct {
		name               string
		expectedStatusCode int
		prepareFunc        func() *http.Request
	}{
		{
			name:               "OK",
			expectedStatusCode: http.StatusCreated,
			prepareFunc: func() *http.Request {
				employeeJSON, err := json.Marshal(models.SignUp{})
				require.NoError(t, err)

				req, err := http.NewRequest(http.MethodPost, handlers.SignUpURL, bytes.NewReader(employeeJSON))
				require.NoError(t, err)

				mockedEmployeeSrv.EXPECT().SignUp(gomock.Any(), gomock.Any()).Return("", "", nil)

				return req
			},
		},
	}

	router := chi.NewRouter()
	router.Post(handlers.SignUpURL, h.SignUp())

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			req := c.prepareFunc()
			recorder := httptest.NewRecorder()

			router.ServeHTTP(recorder, req)

			assert.Equal(t, c.expectedStatusCode, recorder.Code)
		})
	}
}

func TestSignUpErr(t *testing.T) {
	h, mockedEmployeeSrv, ctrl := testHandler(t)
	defer ctrl.Finish()

	cases := []struct {
		name               string
		expectedStatusCode int
		expectedMsg        string
		prepareFunc        func() *http.Request
	}{
		{
			name:               "INVALID REQUEST BODY",
			expectedStatusCode: http.StatusBadRequest,
			expectedMsg:        handlers.ErrInvalidRequestBody.Error(),
			prepareFunc: func() *http.Request {
				req, err := http.NewRequest(http.MethodPost, handlers.SignUpURL, http.NoBody)
				require.NoError(t, err)

				return req
			},
		},
		{
			name:               "VALIDATION ERROR",
			expectedStatusCode: http.StatusBadRequest,
			expectedMsg:        "validation error",
			prepareFunc: func() *http.Request {
				employeeJSON, err := json.Marshal(models.SignUp{})
				require.NoError(t, err)

				req, err := http.NewRequest(http.MethodPost, handlers.SignUpURL, bytes.NewReader(employeeJSON))
				require.NoError(t, err)

				mockedEmployeeSrv.EXPECT().SignUp(gomock.Any(), gomock.Any()).Return("", "", multierror.Append(&multierror.Error{}, errors.New("validation error")))

				return req
			},
		},
		{
			name:               "EMPLOYEE ALREADY EXISTS",
			expectedStatusCode: http.StatusBadRequest,
			expectedMsg:        rep.ErrEmployeeAlreadyExists.Error(),
			prepareFunc: func() *http.Request {
				employeeJSON, err := json.Marshal(models.SignUp{})
				require.NoError(t, err)

				req, err := http.NewRequest(http.MethodPost, handlers.SignUpURL, bytes.NewReader(employeeJSON))
				require.NoError(t, err)

				mockedEmployeeSrv.EXPECT().SignUp(gomock.Any(), gomock.Any()).Return("", "", rep.ErrEmployeeAlreadyExists)

				return req
			},
		},
		{
			name:               "INTERNAL ERROR",
			expectedStatusCode: http.StatusInternalServerError,
			expectedMsg:        handlers.ErrInternal.Error(),
			prepareFunc: func() *http.Request {
				employeeJSON, err := json.Marshal(models.SignUp{})
				require.NoError(t, err)

				req, err := http.NewRequest(http.MethodPost, handlers.SignUpURL, bytes.NewReader(employeeJSON))
				require.NoError(t, err)

				mockedEmployeeSrv.EXPECT().SignUp(gomock.Any(), gomock.Any()).Return("", "", srv.ErrInternal)

				return req
			},
		},
	}

	router := chi.NewRouter()
	router.Post(handlers.SignUpURL, h.SignUp())

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
