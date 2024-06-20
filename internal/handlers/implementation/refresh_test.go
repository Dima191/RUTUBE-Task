package handlerimpl

import (
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
	"testing"
)

func TestRefresh(t *testing.T) {
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
				mockedEmployeeSrv.EXPECT().UpdateTokens(gomock.Any(), gomock.Any()).Return("", "", nil)

				req, err := http.NewRequest(http.MethodPost, handlers.RefreshURL, nil)
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
	router.Post(handlers.RefreshURL, h.Refresh())

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			req := c.prepareFunc()

			recorder := httptest.NewRecorder()

			router.ServeHTTP(recorder, req)

			assert.Equal(t, c.expectedStatusCode, recorder.Code)
		})
	}
}

func TestRefreshErr(t *testing.T) {
	h, mockedEmployeeSrv, ctrl := testHandler(t)
	defer ctrl.Finish()

	cases := []struct {
		name               string
		expectedStatusCode int
		expectedMsg        string
		prepareFunc        func() *http.Request
	}{
		{
			name:               "ERROR NO REFRESH TOKEN",
			expectedStatusCode: http.StatusUnauthorized,
			expectedMsg:        handlers.ErrNoRefreshTokenInCookies.Error(),
			prepareFunc: func() *http.Request {
				req, err := http.NewRequest(http.MethodPost, handlers.RefreshURL, nil)
				require.NoError(t, err)

				return req
			},
		},
		{
			name:               "NO SESSION WITH PROVIDED REFRESH TOKEN",
			expectedStatusCode: http.StatusUnauthorized,
			expectedMsg:        rep.ErrNoSession.Error(),
			prepareFunc: func() *http.Request {
				mockedEmployeeSrv.EXPECT().UpdateTokens(gomock.Any(), gomock.Any()).Return("", "", rep.ErrNoSession)

				req, err := http.NewRequest(http.MethodPost, handlers.RefreshURL, nil)
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
		{
			name:               "REFRESH TOKEN EXPIRED",
			expectedStatusCode: http.StatusUnauthorized,
			expectedMsg:        handlers.ErrRefreshTokenExpired.Error(),
			prepareFunc: func() *http.Request {
				mockedEmployeeSrv.EXPECT().UpdateTokens(gomock.Any(), gomock.Any()).Return("", "", srv.ErrTokenExpired)

				req, err := http.NewRequest(http.MethodPost, handlers.RefreshURL, nil)
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
		{
			name:               "INTERNAL ERROR",
			expectedStatusCode: http.StatusInternalServerError,
			prepareFunc: func() *http.Request {
				mockedEmployeeSrv.EXPECT().UpdateTokens(gomock.Any(), gomock.Any()).Return("", "", srv.ErrInternal)

				req, err := http.NewRequest(http.MethodPost, handlers.RefreshURL, nil)
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
	router.Post(handlers.RefreshURL, h.Refresh())

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			req := c.prepareFunc()

			recorder := httptest.NewRecorder()

			router.ServeHTTP(recorder, req)

			assert.Equal(t, c.expectedStatusCode, recorder.Code)
		})
	}
}
