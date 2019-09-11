// Code generated by MockGen. DO NOT EDIT.
// Source: storage.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	sql "database/sql"
	gomock "github.com/golang/mock/gomock"
	entity "github.com/rafaelsq/boiler/pkg/entity"
	iface "github.com/rafaelsq/boiler/pkg/iface"
	reflect "reflect"
)

// MockStorage is a mock of Storage interface
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// Tx mocks base method
func (m *MockStorage) Tx() (*sql.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tx")
	ret0, _ := ret[0].(*sql.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Tx indicates an expected call of Tx
func (mr *MockStorageMockRecorder) Tx() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tx", reflect.TypeOf((*MockStorage)(nil).Tx))
}

// AddUser mocks base method
func (m *MockStorage) AddUser(ctx context.Context, tx *sql.Tx, name string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", ctx, tx, name)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddUser indicates an expected call of AddUser
func (mr *MockStorageMockRecorder) AddUser(ctx, tx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockStorage)(nil).AddUser), ctx, tx, name)
}

// DeleteUser mocks base method
func (m *MockStorage) DeleteUser(ctx context.Context, tx *sql.Tx, userID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, tx, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser
func (mr *MockStorageMockRecorder) DeleteUser(ctx, tx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockStorage)(nil).DeleteUser), ctx, tx, userID)
}

// FilterUsersID mocks base method
func (m *MockStorage) FilterUsersID(ctx context.Context, filter iface.FilterUsers) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilterUsersID", ctx, filter)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterUsersID indicates an expected call of FilterUsersID
func (mr *MockStorageMockRecorder) FilterUsersID(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterUsersID", reflect.TypeOf((*MockStorage)(nil).FilterUsersID), ctx, filter)
}

// FetchUsers mocks base method
func (m *MockStorage) FetchUsers(ctx context.Context, ID ...int64) ([]*entity.User, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range ID {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FetchUsers", varargs...)
	ret0, _ := ret[0].([]*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchUsers indicates an expected call of FetchUsers
func (mr *MockStorageMockRecorder) FetchUsers(ctx interface{}, ID ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, ID...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchUsers", reflect.TypeOf((*MockStorage)(nil).FetchUsers), varargs...)
}

// AddEmail mocks base method
func (m *MockStorage) AddEmail(ctx context.Context, tx *sql.Tx, userID int64, address string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddEmail", ctx, tx, userID, address)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddEmail indicates an expected call of AddEmail
func (mr *MockStorageMockRecorder) AddEmail(ctx, tx, userID, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEmail", reflect.TypeOf((*MockStorage)(nil).AddEmail), ctx, tx, userID, address)
}

// DeleteEmail mocks base method
func (m *MockStorage) DeleteEmail(ctx context.Context, tx *sql.Tx, emailID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEmail", ctx, tx, emailID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEmail indicates an expected call of DeleteEmail
func (mr *MockStorageMockRecorder) DeleteEmail(ctx, tx, emailID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEmail", reflect.TypeOf((*MockStorage)(nil).DeleteEmail), ctx, tx, emailID)
}

// DeleteEmailsByUserID mocks base method
func (m *MockStorage) DeleteEmailsByUserID(ctx context.Context, tx *sql.Tx, userID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEmailsByUserID", ctx, tx, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEmailsByUserID indicates an expected call of DeleteEmailsByUserID
func (mr *MockStorageMockRecorder) DeleteEmailsByUserID(ctx, tx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEmailsByUserID", reflect.TypeOf((*MockStorage)(nil).DeleteEmailsByUserID), ctx, tx, userID)
}

// FilterEmails mocks base method
func (m *MockStorage) FilterEmails(ctx context.Context, filter iface.FilterEmails) ([]*entity.Email, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilterEmails", ctx, filter)
	ret0, _ := ret[0].([]*entity.Email)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterEmails indicates an expected call of FilterEmails
func (mr *MockStorageMockRecorder) FilterEmails(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterEmails", reflect.TypeOf((*MockStorage)(nil).FilterEmails), ctx, filter)
}
