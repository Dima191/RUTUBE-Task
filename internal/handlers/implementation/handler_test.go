package handlerimpl

import (
	"github.com/Dima191/RUTUBE-Task/internal/handlers"
	employeesrv "github.com/Dima191/RUTUBE-Task/internal/service/employee"
	stubwriter "github.com/Dima191/RUTUBE-Task/pkg/stub_writer"
	tokenmanager "github.com/Dima191/RUTUBE-Task/pkg/token_manager"
	"go.uber.org/mock/gomock"
	"log/slog"
	"testing"
)

func testHandler(t *testing.T) (h handlers.Handler, mockedEmployeeSrv *employeesrv.MockService, ctrl *gomock.Controller) {
	t.Helper()

	ctrl = gomock.NewController(t)

	mockedEmployeeSrv = employeesrv.NewMockService(ctrl)
	mockedTokenManager := tokenmanager.NewMockTokenManager(ctrl)

	logger := slog.New(slog.NewTextHandler(&stubwriter.Writer{}, nil))

	h = New(mockedEmployeeSrv, mockedTokenManager, logger)

	return h, mockedEmployeeSrv, ctrl
}
