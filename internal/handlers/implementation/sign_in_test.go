package handlerimpl

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/Dima191/RUTUBE-Task/internal/handlers"
	"github.com/Dima191/RUTUBE-Task/internal/models"
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

func TestSignIn(t *testing.T) {
	h, mockedEmployeeSrv, ctrl := testHandler(t)
	defer ctrl.Finish()

	cases := []struct {
		name               string
		expectedStatusCode int
		prepareFunc        func() *http.Request
	}{
		{
			name:               "OK",
			expectedStatusCode: http.StatusOK,
			prepareFunc: func() *http.Request {
				mockedEmployeeSrv.EXPECT().SignIn(gomock.Any(), gomock.Any()).Return("", "", nil)

				employeeJSON, err := json.Marshal(models.SignIn{})
				require.NoError(t, err)

				req, err := http.NewRequest(http.MethodPost, handlers.SignInURL, bytes.NewReader(employeeJSON))
				require.NoError(t, err)

				return req
			},
		},
	}

	router := chi.NewRouter()
	router.Post(handlers.SignInURL, h.SignIn())

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			req := c.prepareFunc()

			recorder := httptest.NewRecorder()

			router.ServeHTTP(recorder, req)

			assert.Equal(t, c.expectedStatusCode, recorder.Code)
		})
	}
}

func TestSignInErr(t *testing.T) {
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
				req, err := http.NewRequest(http.MethodPost, handlers.SignInURL, http.NoBody)
				require.NoError(t, err)

				return req
			},
		},
		{
			name:               "INVALID CREDENTIALS",
			expectedStatusCode: http.StatusBadRequest,
			expectedMsg:        srv.ErrInvalidLoginOrPassword.Error(),
			prepareFunc: func() *http.Request {
				mockedEmployeeSrv.EXPECT().SignIn(gomock.Any(), gomock.Any()).Return("", "", srv.ErrInvalidLoginOrPassword)

				employeeJSON, err := json.Marshal(models.SignIn{})
				require.NoError(t, err)

				req, err := http.NewRequest(http.MethodPost, handlers.SignInURL, bytes.NewReader(employeeJSON))
				require.NoError(t, err)

				return req
			},
		},
		{
			name:               "VALIDATION ERROR",
			expectedStatusCode: http.StatusBadRequest,
			expectedMsg:        "validation error",
			prepareFunc: func() *http.Request {
				mockedEmployeeSrv.EXPECT().SignIn(gomock.Any(), gomock.Any()).Return("", "", multierror.Append(&multierror.Error{}, errors.New("validation error")))

				employeeJSON, err := json.Marshal(models.SignIn{})
				require.NoError(t, err)

				req, err := http.NewRequest(http.MethodPost, handlers.SignInURL, bytes.NewReader(employeeJSON))
				require.NoError(t, err)

				return req
			},
		},
		{
			name:               "INTERNAL ERROR",
			expectedStatusCode: http.StatusInternalServerError,
			expectedMsg:        handlers.ErrInternal.Error(),
			prepareFunc: func() *http.Request {
				mockedEmployeeSrv.EXPECT().SignIn(gomock.Any(), gomock.Any()).Return("", "", srv.ErrInternal)

				employeeJSON, err := json.Marshal(models.SignIn{})
				require.NoError(t, err)

				req, err := http.NewRequest(http.MethodPost, handlers.SignInURL, bytes.NewReader(employeeJSON))
				require.NoError(t, err)

				return req
			},
		},
	}

	router := chi.NewRouter()
	router.Post(handlers.SignInURL, h.SignIn())

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
