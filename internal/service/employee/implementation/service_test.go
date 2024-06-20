package employeesrvimpl

import (
	rep "github.com/Dima191/RUTUBE-Task/internal/repository"
	employeesrv "github.com/Dima191/RUTUBE-Task/internal/service/employee"
	stubwriter "github.com/Dima191/RUTUBE-Task/pkg/stub_writer"
	"github.com/Dima191/RUTUBE-Task/pkg/token_manager"
	"go.uber.org/mock/gomock"
	"log/slog"
	"testing"
)

func testService(t *testing.T) (employeesrv.Service, *rep.MockRepository, *token_manager.MockTokenManager, *gomock.Controller) {
	t.Helper()

	ctrl := gomock.NewController(t)

	mockedRep := rep.NewMockRepository(ctrl)
	mockedTokenManager := token_manager.NewMockTokenManager(ctrl)

	logger := slog.New(slog.NewTextHandler(&stubwriter.Writer{}, &slog.HandlerOptions{}))

	return New(mockedRep, mockedTokenManager, logger), mockedRep, mockedTokenManager, ctrl
}
