package handlerimpl

import (
	"github.com/Dima191/RUTUBE-Task/internal/handlers"
	"github.com/Dima191/RUTUBE-Task/internal/middlewares"
	srv "github.com/Dima191/RUTUBE-Task/internal/service/employee"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogOut(t *testing.T) {
	h, mockedEmployeeSrv, ctrl := testHandler(t)
	defer ctrl.Finish()

	cases := []struct {
		name                string
		expectedStatusCodes int
		prepareFunc         func() (req *http.Request)
	}{
		{
			name:                "OK",
			expectedStatusCodes: http.StatusOK,
			prepareFunc: func() (req *http.Request) {
				mockedEmployeeSrv.EXPECT().LogOut(gomock.Any(), gomock.Any()).Return(nil)

				req, err := http.NewRequest(http.MethodPost, handlers.LogOutURL, nil)
				require.NoError(t, err)

				req.AddCookie(&http.Cookie{
					Name:     middlewares.RefreshTokenCookiesKey,
					Path:     "/",
					HttpOnly: true,
					Secure:   true,
					SameSite: http.SameSiteStrictMode,
				})

				return req
			},
		},
	}

	router := chi.NewRouter()
	router.Post(handlers.LogOutURL, h.LogOut())

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			req := c.prepareFunc()

			recorder := httptest.NewRecorder()

			router.ServeHTTP(recorder, req)

			require.Equal(t, c.expectedStatusCodes, recorder.Code)
		})
	}
}

func TestLogOutErr(t *testing.T) {
	h, mockedEmployeeSrv, ctrl := testHandler(t)
	defer ctrl.Finish()

	cases := []struct {
		name                string
		expectedStatusCodes int
		expectedMsg         string
		prepareFunc         func() (req *http.Request)
	}{
		{
			name:                "ERROR NO REFRESH TOKEN",
			expectedStatusCodes: http.StatusUnauthorized,
			expectedMsg:         handlers.ErrNoRefreshTokenInCookies.Error(),
			prepareFunc: func() (req *http.Request) {
				req, err := http.NewRequest(http.MethodPost, handlers.LogOutURL, nil)
				require.NoError(t, err)

				return req
			},
		},
		{
			name:                "INTERNAL ERROR",
			expectedStatusCodes: http.StatusInternalServerError,
			prepareFunc: func() (req *http.Request) {
				mockedEmployeeSrv.EXPECT().LogOut(gomock.Any(), gomock.Any()).Return(srv.ErrInternal)

				req, err := http.NewRequest(http.MethodPost, handlers.LogOutURL, nil)
				require.NoError(t, err)

				req.AddCookie(&http.Cookie{
					Name:     middlewares.RefreshTokenCookiesKey,
					Path:     "/",
					HttpOnly: true,
					Secure:   true,
					SameSite: http.SameSiteStrictMode,
				})

				return req
			},
		},
	}

	router := chi.NewRouter()
	router.Post(handlers.LogOutURL, h.LogOut())

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			req := c.prepareFunc()

			recorder := httptest.NewRecorder()

			router.ServeHTTP(recorder, req)

			require.Equal(t, c.expectedStatusCodes, recorder.Code)
		})
	}
}
