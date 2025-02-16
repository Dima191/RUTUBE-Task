// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/token_manager/token_manager.go
//
// Generated by this command:
//
//	mockgen -source pkg/token_manager/token_manager.go -destination pkg/token_manager/mock_token_manager.go
//

// Package mock_tokenmanager is a generated GoMock package.
package token_manager

import (
	reflect "reflect"

	models "github.com/Dima191/RUTUBE-Task/internal/models"
	jwt "github.com/golang-jwt/jwt/v5"
	gomock "go.uber.org/mock/gomock"
)

// MockTokenManager is a mock of TokenManager interface.
type MockTokenManager struct {
	ctrl     *gomock.Controller
	recorder *MockTokenManagerMockRecorder
}

// MockTokenManagerMockRecorder is the mock recorder for MockTokenManager.
type MockTokenManagerMockRecorder struct {
	mock *MockTokenManager
}

// NewMockTokenManager creates a new mock instance.
func NewMockTokenManager(ctrl *gomock.Controller) *MockTokenManager {
	mock := &MockTokenManager{ctrl: ctrl}
	mock.recorder = &MockTokenManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenManager) EXPECT() *MockTokenManagerMockRecorder {
	return m.recorder
}

// GenerateAccessToken mocks base method.
func (m *MockTokenManager) GenerateAccessToken(employeeID uint32) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateAccessToken", employeeID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateAccessToken indicates an expected call of GenerateAccessToken.
func (mr *MockTokenManagerMockRecorder) GenerateAccessToken(employeeID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateAccessToken", reflect.TypeOf((*MockTokenManager)(nil).GenerateAccessToken), employeeID)
}

// GenerateRefreshToken mocks base method.
func (m *MockTokenManager) GenerateRefreshToken() (models.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateRefreshToken")
	ret0, _ := ret[0].(models.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateRefreshToken indicates an expected call of GenerateRefreshToken.
func (mr *MockTokenManagerMockRecorder) GenerateRefreshToken() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateRefreshToken", reflect.TypeOf((*MockTokenManager)(nil).GenerateRefreshToken))
}

// Parse mocks base method.
func (m *MockTokenManager) Parse(accessToken string) (jwt.Claims, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parse", accessToken)
	ret0, _ := ret[0].(jwt.Claims)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Parse indicates an expected call of Parse.
func (mr *MockTokenManagerMockRecorder) Parse(accessToken any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parse", reflect.TypeOf((*MockTokenManager)(nil).Parse), accessToken)
}
