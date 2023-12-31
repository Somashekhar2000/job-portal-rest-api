// Code generated by MockGen. DO NOT EDIT.
// Source: companyRepository.go
//
// Generated by this command:
//
//	mockgen -source=companyRepository.go -destination=companyRepository_mock.go -package=repository
//
// Package repository is a generated GoMock package.
package repository

import (
	model "job-portal-api/internal/model"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockComapnyRepo is a mock of ComapnyRepo interface.
type MockComapnyRepo struct {
	ctrl     *gomock.Controller
	recorder *MockComapnyRepoMockRecorder
}

// MockComapnyRepoMockRecorder is the mock recorder for MockComapnyRepo.
type MockComapnyRepoMockRecorder struct {
	mock *MockComapnyRepo
}

// NewMockComapnyRepo creates a new mock instance.
func NewMockComapnyRepo(ctrl *gomock.Controller) *MockComapnyRepo {
	mock := &MockComapnyRepo{ctrl: ctrl}
	mock.recorder = &MockComapnyRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockComapnyRepo) EXPECT() *MockComapnyRepoMockRecorder {
	return m.recorder
}

// CreateComapny mocks base method.
func (m *MockComapnyRepo) CreateComapny(company model.Company) (model.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateComapny", company)
	ret0, _ := ret[0].(model.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateComapny indicates an expected call of CreateComapny.
func (mr *MockComapnyRepoMockRecorder) CreateComapny(company any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComapny", reflect.TypeOf((*MockComapnyRepo)(nil).CreateComapny), company)
}

// GetAllCompanies mocks base method.
func (m *MockComapnyRepo) GetAllCompanies() ([]model.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCompanies")
	ret0, _ := ret[0].([]model.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCompanies indicates an expected call of GetAllCompanies.
func (mr *MockComapnyRepoMockRecorder) GetAllCompanies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCompanies", reflect.TypeOf((*MockComapnyRepo)(nil).GetAllCompanies))
}

// GetCompanyByID mocks base method.
func (m *MockComapnyRepo) GetCompanyByID(cID uint64) (model.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCompanyByID", cID)
	ret0, _ := ret[0].(model.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCompanyByID indicates an expected call of GetCompanyByID.
func (mr *MockComapnyRepoMockRecorder) GetCompanyByID(cID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCompanyByID", reflect.TypeOf((*MockComapnyRepo)(nil).GetCompanyByID), cID)
}
