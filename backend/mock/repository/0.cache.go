// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/repository/0.cache.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"
	time "time"

	repository "github.com/cocoide/tech-guide/pkg/repository"
	gomock "github.com/golang/mock/gomock"
)

// MockCacheRepo is a mock of CacheRepo interface.
type MockCacheRepo struct {
	ctrl     *gomock.Controller
	recorder *MockCacheRepoMockRecorder
}

// MockCacheRepoMockRecorder is the mock recorder for MockCacheRepo.
type MockCacheRepoMockRecorder struct {
	mock *MockCacheRepo
}

// NewMockCacheRepo creates a new mock instance.
func NewMockCacheRepo(ctrl *gomock.Controller) *MockCacheRepo {
	mock := &MockCacheRepo{ctrl: ctrl}
	mock.recorder = &MockCacheRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCacheRepo) EXPECT() *MockCacheRepoMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockCacheRepo) Delete(key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCacheRepoMockRecorder) Delete(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCacheRepo)(nil).Delete), key)
}

// ExtendExpiry mocks base method.
func (m *MockCacheRepo) ExtendExpiry(key string, extension time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtendExpiry", key, extension)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExtendExpiry indicates an expected call of ExtendExpiry.
func (mr *MockCacheRepoMockRecorder) ExtendExpiry(key, extension interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtendExpiry", reflect.TypeOf((*MockCacheRepo)(nil).ExtendExpiry), key, extension)
}

// Get mocks base method.
func (m *MockCacheRepo) Get(key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCacheRepoMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCacheRepo)(nil).Get), key)
}

// Set mocks base method.
func (m *MockCacheRepo) Set(key, value string, expire time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", key, value, expire)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockCacheRepoMockRecorder) Set(key, value, expire interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockCacheRepo)(nil).Set), key, value, expire)
}

// Update mocks base method.
func (m *MockCacheRepo) Update(key, value string, expire time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", key, value, expire)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockCacheRepoMockRecorder) Update(key, value, expire interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCacheRepo)(nil).Update), key, value, expire)
}

// WithCtx mocks base method.
func (m *MockCacheRepo) WithCtx(ctx context.Context) repository.CacheRepo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithCtx", ctx)
	ret0, _ := ret[0].(repository.CacheRepo)
	return ret0
}

// WithCtx indicates an expected call of WithCtx.
func (mr *MockCacheRepoMockRecorder) WithCtx(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithCtx", reflect.TypeOf((*MockCacheRepo)(nil).WithCtx), ctx)
}
