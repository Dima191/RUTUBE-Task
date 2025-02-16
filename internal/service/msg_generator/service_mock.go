// Code generated by MockGen. DO NOT EDIT.
// Source: internal/service/msg_generator/service.go
//
// Generated by this command:
//
//	mockgen -source internal/service/msg_generator/service.go -destination internal/service/msg_generator/service_mock.go
//

// Package mock_msggenerator is a generated GoMock package.
package msggenerator

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// Generate mocks base method.
func (m *MockService) Generate(ctx context.Context, subscriberFullName, celebrantFullName string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generate", ctx, subscriberFullName, celebrantFullName)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Generate indicates an expected call of Generate.
func (mr *MockServiceMockRecorder) Generate(ctx, subscriberFullName, celebrantFullName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generate", reflect.TypeOf((*MockService)(nil).Generate), ctx, subscriberFullName, celebrantFullName)
}
