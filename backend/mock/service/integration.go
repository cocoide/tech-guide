// Code generated by MockGen. DO NOT EDIT.
// Source: 0.integration.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	model "github.com/cocoide/tech-guide/pkg/domain/model"
	dto "github.com/cocoide/tech-guide/pkg/domain/model/dto"
	gomock "github.com/golang/mock/gomock"
)

// MockOpenAIService is a mock of OpenAIService interface.
type MockOpenAIService struct {
	ctrl     *gomock.Controller
	recorder *MockOpenAIServiceMockRecorder
}

// MockOpenAIServiceMockRecorder is the mock recorder for MockOpenAIService.
type MockOpenAIServiceMockRecorder struct {
	mock *MockOpenAIService
}

// NewMockOpenAIService creates a new mock instance.
func NewMockOpenAIService(ctrl *gomock.Controller) *MockOpenAIService {
	mock := &MockOpenAIService{ctrl: ctrl}
	mock.recorder = &MockOpenAIServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOpenAIService) EXPECT() *MockOpenAIServiceMockRecorder {
	return m.recorder
}

// AsyncGetAnswerFromPrompt mocks base method.
func (m *MockOpenAIService) AsyncGetAnswerFromPrompt(prompt string, variability float32) <-chan string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AsyncGetAnswerFromPrompt", prompt, variability)
	ret0, _ := ret[0].(<-chan string)
	return ret0
}

// AsyncGetAnswerFromPrompt indicates an expected call of AsyncGetAnswerFromPrompt.
func (mr *MockOpenAIServiceMockRecorder) AsyncGetAnswerFromPrompt(prompt, variability interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AsyncGetAnswerFromPrompt", reflect.TypeOf((*MockOpenAIService)(nil).AsyncGetAnswerFromPrompt), prompt, variability)
}

// GetAnswerFromPrompt mocks base method.
func (m *MockOpenAIService) GetAnswerFromPrompt(prompt string, variability float32) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAnswerFromPrompt", prompt, variability)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAnswerFromPrompt indicates an expected call of GetAnswerFromPrompt.
func (mr *MockOpenAIServiceMockRecorder) GetAnswerFromPrompt(prompt, variability interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAnswerFromPrompt", reflect.TypeOf((*MockOpenAIService)(nil).GetAnswerFromPrompt), prompt, variability)
}

// MockOGPService is a mock of OGPService interface.
type MockOGPService struct {
	ctrl     *gomock.Controller
	recorder *MockOGPServiceMockRecorder
}

// MockOGPServiceMockRecorder is the mock recorder for MockOGPService.
type MockOGPServiceMockRecorder struct {
	mock *MockOGPService
}

// NewMockOGPService creates a new mock instance.
func NewMockOGPService(ctrl *gomock.Controller) *MockOGPService {
	mock := &MockOGPService{ctrl: ctrl}
	mock.recorder = &MockOGPServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOGPService) EXPECT() *MockOGPServiceMockRecorder {
	return m.recorder
}

// GetOGPByURL mocks base method.
func (m *MockOGPService) GetOGPByURL(url string) (*dto.OGPResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOGPByURL", url)
	ret0, _ := ret[0].(*dto.OGPResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOGPByURL indicates an expected call of GetOGPByURL.
func (mr *MockOGPServiceMockRecorder) GetOGPByURL(url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOGPByURL", reflect.TypeOf((*MockOGPService)(nil).GetOGPByURL), url)
}

// MockTechFeedService is a mock of TechFeedService interface.
type MockTechFeedService struct {
	ctrl     *gomock.Controller
	recorder *MockTechFeedServiceMockRecorder
}

// MockTechFeedServiceMockRecorder is the mock recorder for MockTechFeedService.
type MockTechFeedServiceMockRecorder struct {
	mock *MockTechFeedService
}

// NewMockTechFeedService creates a new mock instance.
func NewMockTechFeedService(ctrl *gomock.Controller) *MockTechFeedService {
	mock := &MockTechFeedService{ctrl: ctrl}
	mock.recorder = &MockTechFeedServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTechFeedService) EXPECT() *MockTechFeedServiceMockRecorder {
	return m.recorder
}

// GetQiitaTrendFeed mocks base method.
func (m *MockTechFeedService) GetQiitaTrendFeed(limit, save int, start string) ([]*dto.QiitaArticleAPI, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQiitaTrendFeed", limit, save, start)
	ret0, _ := ret[0].([]*dto.QiitaArticleAPI)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQiitaTrendFeed indicates an expected call of GetQiitaTrendFeed.
func (mr *MockTechFeedServiceMockRecorder) GetQiitaTrendFeed(limit, save, start interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQiitaTrendFeed", reflect.TypeOf((*MockTechFeedService)(nil).GetQiitaTrendFeed), limit, save, start)
}

// GetZennTrendFeed mocks base method.
func (m *MockTechFeedService) GetZennTrendFeed() ([]model.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetZennTrendFeed")
	ret0, _ := ret[0].([]model.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetZennTrendFeed indicates an expected call of GetZennTrendFeed.
func (mr *MockTechFeedServiceMockRecorder) GetZennTrendFeed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetZennTrendFeed", reflect.TypeOf((*MockTechFeedService)(nil).GetZennTrendFeed))
}
