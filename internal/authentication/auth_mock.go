// Code generated by MockGen. DO NOT EDIT.
// Source: auth.go
//
// Generated by this command:
//
//	mockgen -source=auth.go -destination=auth_mock.go -package=service
//
// Package service is a generated GoMock package.
package authentication

import (
	reflect "reflect"

	jwt "github.com/golang-jwt/jwt/v5"
	gomock "go.uber.org/mock/gomock"
)

// MockAuthenticaton is a mock of Authenticaton interface.
type MockAuthenticaton struct {
	ctrl     *gomock.Controller
	recorder *MockAuthenticatonMockRecorder
}

// MockAuthenticatonMockRecorder is the mock recorder for MockAuthenticaton.
type MockAuthenticatonMockRecorder struct {
	mock *MockAuthenticaton
}


// NewMockAuthenticaton creates a new mock instance.
func NewMockAuthenticaton(ctrl *gomock.Controller) *MockAuthenticaton {
	mock := &MockAuthenticaton{ctrl: ctrl}
	mock.recorder = &MockAuthenticatonMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthenticaton) EXPECT() *MockAuthenticatonMockRecorder {
	return m.recorder
}

// GenerateToken mocks base method.
func (m *MockAuthenticaton) GenerateToken(claims jwt.RegisteredClaims) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", claims)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockAuthenticatonMockRecorder) GenerateToken(claims any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockAuthenticaton)(nil).GenerateToken), claims)
}

// ValidateToken mocks base method.
func (m *MockAuthenticaton) ValidateToken(token string) (jwt.RegisteredClaims, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateToken", token)
	ret0, _ := ret[0].(jwt.RegisteredClaims)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateToken indicates an expected call of ValidateToken.
func (mr *MockAuthenticatonMockRecorder) ValidateToken(token any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateToken", reflect.TypeOf((*MockAuthenticaton)(nil).ValidateToken), token)
}
