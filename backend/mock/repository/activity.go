// Code generated by MockGen. DO NOT EDIT.
// Source: activity.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	model "github.com/cocoide/tech-guide/pkg/domain/model"
	gomock "github.com/golang/mock/gomock"
)

// MockActivityRepo is a mock of ActivityRepo interface.
type MockActivityRepo struct {
	ctrl     *gomock.Controller
	recorder *MockActivityRepoMockRecorder
}

// MockActivityRepoMockRecorder is the mock recorder for MockActivityRepo.
type MockActivityRepoMockRecorder struct {
	mock *MockActivityRepo
}

// NewMockActivityRepo creates a new mock instance.
func NewMockActivityRepo(ctrl *gomock.Controller) *MockActivityRepo {
	mock := &MockActivityRepo{ctrl: ctrl}
	mock.recorder = &MockActivityRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockActivityRepo) EXPECT() *MockActivityRepoMockRecorder {
	return m.recorder
}

// BatchCreateContributions mocks base method.
func (m *MockActivityRepo) BatchCreateContributions(contribute []*model.Contribution) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchCreateContributions", contribute)
	ret0, _ := ret[0].(error)
	return ret0
}

// BatchCreateContributions indicates an expected call of BatchCreateContributions.
func (mr *MockActivityRepoMockRecorder) BatchCreateContributions(contribute interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchCreateContributions", reflect.TypeOf((*MockActivityRepo)(nil).BatchCreateContributions), contribute)
}

// CreateAchievement mocks base method.
func (m *MockActivityRepo) CreateAchievement(achieve *model.AccountsToAchievements) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAchievement", achieve)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAchievement indicates an expected call of CreateAchievement.
func (mr *MockActivityRepoMockRecorder) CreateAchievement(achieve interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAchievement", reflect.TypeOf((*MockActivityRepo)(nil).CreateAchievement), achieve)
}

// CreateContribution mocks base method.
func (m *MockActivityRepo) CreateContribution(contribute *model.Contribution) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateContribution", contribute)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateContribution indicates an expected call of CreateContribution.
func (mr *MockActivityRepoMockRecorder) CreateContribution(contribute interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateContribution", reflect.TypeOf((*MockActivityRepo)(nil).CreateContribution), contribute)
}

// GetContributionsByAccountID mocks base method.
func (m *MockActivityRepo) GetContributionsByAccountID(accountID int) ([]*model.Contribution, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContributionsByAccountID", accountID)
	ret0, _ := ret[0].([]*model.Contribution)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContributionsByAccountID indicates an expected call of GetContributionsByAccountID.
func (mr *MockActivityRepoMockRecorder) GetContributionsByAccountID(accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContributionsByAccountID", reflect.TypeOf((*MockActivityRepo)(nil).GetContributionsByAccountID), accountID)
}

// UpdateContribution mocks base method.
func (m *MockActivityRepo) UpdateContribution(contribute *model.Contribution) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateContribution", contribute)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateContribution indicates an expected call of UpdateContribution.
func (mr *MockActivityRepoMockRecorder) UpdateContribution(contribute interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateContribution", reflect.TypeOf((*MockActivityRepo)(nil).UpdateContribution), contribute)
}