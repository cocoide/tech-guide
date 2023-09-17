// Code generated by MockGen. DO NOT EDIT.
// Source: topic.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	"context"
	reflect "reflect"
	time "time"

	model "github.com/cocoide/tech-guide/pkg/domain/model"
	gomock "github.com/golang/mock/gomock"
)

// MockTopicRepo is a mock of TopicRepo interface.
type MockTopicRepo struct {
	ctrl     *gomock.Controller
	recorder *MockTopicRepoMockRecorder
}

// BatchCreate implements repository.Repository.
func (*MockTopicRepo) BatchCreate(articles []*model.Article) ([]int, error) {
	panic("unimplemented")
}

// BatchCreateContributions implements repository.Repository.
func (*MockTopicRepo) BatchCreateContributions(contribute []*model.Contribution) error {
	panic("unimplemented")
}

// BatchGetArticlesByTopicIDsAndSourceID implements repository.Repository.
func (*MockTopicRepo) BatchGetArticlesByTopicIDsAndSourceID(topicIDs []int, sourceIDs []int, pageIndex int, pageSize int) ([]model.Article, error) {
	panic("unimplemented")
}

// CheckArticleExistsByURL implements repository.Repository.
func (*MockTopicRepo) CheckArticleExistsByURL(url string) (bool, error) {
	panic("unimplemented")
}

// CheckExistByEmail implements repository.Repository.
func (*MockTopicRepo) CheckExistByEmail(email string) (bool, error) {
	panic("unimplemented")
}

// CountFavoritesByArticleID implements repository.Repository.
func (*MockTopicRepo) CountFavoritesByArticleID(articleID int) (int64, error) {
	panic("unimplemented")
}

// CreateAccount implements repository.Repository.
func (*MockTopicRepo) CreateAccount(account *model.Account) (*model.Account, error) {
	panic("unimplemented")
}

// CreateAchievement implements repository.Repository.
func (*MockTopicRepo) CreateAchievement(achieve *model.AccountsToAchievements) error {
	panic("unimplemented")
}

// CreateArticle implements repository.Repository.
func (*MockTopicRepo) CreateArticle(article *model.Article) (int, error) {
	panic("unimplemented")
}

// CreateBookmark implements repository.Repository.
func (*MockTopicRepo) CreateBookmark(bookmark *model.Bookmark) error {
	panic("unimplemented")
}

// CreateCollection implements repository.Repository.
func (*MockTopicRepo) CreateCollection(collection *model.Collection) error {
	panic("unimplemented")
}

// CreateCollectionWithBookmark implements repository.Repository.
func (*MockTopicRepo) CreateCollectionWithBookmark(collectino *model.Collection, articleId int) error {
	panic("unimplemented")
}

// CreateComment implements repository.Repository.
func (*MockTopicRepo) CreateComment(comment *model.Comment) error {
	panic("unimplemented")
}

// CreateContribution implements repository.Repository.
func (*MockTopicRepo) CreateContribution(contribute *model.Contribution) error {
	panic("unimplemented")
}

// CreateTopicToArticle implements repository.Repository.
func (*MockTopicRepo) CreateTopicToArticle(topicToArticles []model.TopicsToArticles) error {
	panic("unimplemented")
}

// DoFavoriteArticle implements repository.Repository.
func (*MockTopicRepo) DoFavoriteArticle(articleID int, accountID int) error {
	panic("unimplemented")
}

// DoFollowSource implements repository.Repository.
func (*MockTopicRepo) DoFollowSource(accountID int, souceID int) error {
	panic("unimplemented")
}

// DoInTx implements repository.Repository.
func (*MockTopicRepo) DoInTx(ctx context.Context, fn func(ctx context.Context) error) error {
	panic("unimplemented")
}

// FindIDByDomain implements repository.Repository.
func (*MockTopicRepo) FindIDByDomain(domain string) (int, error) {
	panic("unimplemented")
}

// GetAccountProfile implements repository.Repository.
func (*MockTopicRepo) GetAccountProfile(id int) (*model.Account, error) {
	panic("unimplemented")
}

// GetAllAccountIDs implements repository.Repository.
func (*MockTopicRepo) GetAllAccountIDs() ([]int, error) {
	panic("unimplemented")
}

// GetAllSources implements repository.Repository.
func (*MockTopicRepo) GetAllSources() ([]model.Source, error) {
	panic("unimplemented")
}

// GetArticleByID implements repository.Repository.
func (*MockTopicRepo) GetArticleByID(articleID int) (*model.Article, error) {
	panic("unimplemented")
}

// GetArticleIDByURL implements repository.Repository.
func (*MockTopicRepo) GetArticleIDByURL(url string) (int, error) {
	panic("unimplemented")
}

// GetArticleWithRelatedDataByID implements repository.Repository.
func (*MockTopicRepo) GetArticleWithRelatedDataByID(articleID int) (*model.Article, error) {
	panic("unimplemented")
}

// GetArticlesByIDs implements repository.Repository.
func (*MockTopicRepo) GetArticlesByIDs(articleIDs []int) ([]model.Article, error) {
	panic("unimplemented")
}

// GetArticlesBySourceID implements repository.Repository.
func (*MockTopicRepo) GetArticlesBySourceID(sourceID int, pageIndex int, pageSize int) ([]model.Article, error) {
	panic("unimplemented")
}

// GetArticlesByTopicID implements repository.Repository.
func (*MockTopicRepo) GetArticlesByTopicID(topicID int, pageIndex int, pageSize int) ([]model.Article, error) {
	panic("unimplemented")
}

// GetArticlesByTopicIDs implements repository.Repository.
func (*MockTopicRepo) GetArticlesByTopicIDs(topicIDs []int, omitArticleId int) ([]model.TopicsToArticles, error) {
	panic("unimplemented")
}

// GetByEmail implements repository.Repository.
func (*MockTopicRepo) GetByEmail(email string) (*model.Account, error) {
	panic("unimplemented")
}

// GetCollectionAuthorID implements repository.Repository.
func (*MockTopicRepo) GetCollectionAuthorID(collectionId int) (int, error) {
	panic("unimplemented")
}

// GetCollectionByID implements repository.Repository.
func (*MockTopicRepo) GetCollectionByID(id int) (*model.Collection, error) {
	panic("unimplemented")
}

// GetCollectionsByAccountID implements repository.Repository.
func (*MockTopicRepo) GetCollectionsByAccountID(accountId int) ([]*model.Collection, error) {
	panic("unimplemented")
}

// GetComments implements repository.Repository.
func (*MockTopicRepo) GetComments(articleID int) ([]model.Comment, error) {
	panic("unimplemented")
}

// GetContributionsByAccountID implements repository.Repository.
func (*MockTopicRepo) GetContributionsByAccountID(accountID int) ([]*model.Contribution, error) {
	panic("unimplemented")
}

// GetFollowingSourceIDs implements repository.Repository.
func (*MockTopicRepo) GetFollowingSourceIDs(accountID int) ([]int, error) {
	panic("unimplemented")
}

// GetFollowingSources implements repository.Repository.
func (*MockTopicRepo) GetFollowingSources(accountId int) ([]model.Source, error) {
	panic("unimplemented")
}

// GetLatestArticleByLimitWithSourceData implements repository.Repository.
func (*MockTopicRepo) GetLatestArticleByLimitWithSourceData(pageIndex int, pageSize int) ([]*model.Article, error) {
	panic("unimplemented")
}

// GetPopularSources implements repository.Repository.
func (*MockTopicRepo) GetPopularSources(limit int) ([]model.Source, error) {
	panic("unimplemented")
}

// GetSourceData implements repository.Repository.
func (*MockTopicRepo) GetSourceData(sourceID int) (*model.Source, error) {
	panic("unimplemented")
}

// GetTagsAndWeightsByArticleID implements repository.Repository.
func (*MockTopicRepo) GetTagsAndWeightsByArticleID(articleID int) ([]model.TopicsToArticles, error) {
	panic("unimplemented")
}

// GetTopicsByID implements repository.Repository.
func (*MockTopicRepo) GetTopicsByID(articleId int) ([]model.Topic, error) {
	panic("unimplemented")
}

// IsFollowingSource implements repository.Repository.
func (*MockTopicRepo) IsFollowingSource(accountID int, sourceID int) (bool, error) {
	panic("unimplemented")
}

// UnFavoriteArticle implements repository.Repository.
func (*MockTopicRepo) UnFavoriteArticle(articleID int, accountID int) error {
	panic("unimplemented")
}

// UnFollowSource implements repository.Repository.
func (*MockTopicRepo) UnFollowSource(accountID int, sourceID int) error {
	panic("unimplemented")
}

// UpdateContribution implements repository.Repository.
func (*MockTopicRepo) UpdateContribution(contribute *model.Contribution) error {
	panic("unimplemented")
}

// UpdateSummaryByID implements repository.Repository.
func (*MockTopicRepo) UpdateSummaryByID(id int, summary string) error {
	panic("unimplemented")
}

// MockTopicRepoMockRecorder is the mock recorder for MockTopicRepo.
type MockTopicRepoMockRecorder struct {
	mock *MockTopicRepo
}

// NewMockTopicRepo creates a new mock instance.
func NewMockTopicRepo(ctrl *gomock.Controller) *MockTopicRepo {
	mock := &MockTopicRepo{ctrl: ctrl}
	mock.recorder = &MockTopicRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTopicRepo) EXPECT() *MockTopicRepoMockRecorder {
	return m.recorder
}

// CreateTopics mocks base method.
func (m *MockTopicRepo) CreateTopics(topics []model.Topic) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTopics", topics)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTopics indicates an expected call of CreateTopics.
func (mr *MockTopicRepoMockRecorder) CreateTopics(topics interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTopics", reflect.TypeOf((*MockTopicRepo)(nil).CreateTopics), topics)
}

// DoFollowTopic mocks base method.
func (m *MockTopicRepo) DoFollowTopic(accountID, topicID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoFollowTopic", accountID, topicID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DoFollowTopic indicates an expected call of DoFollowTopic.
func (mr *MockTopicRepoMockRecorder) DoFollowTopic(accountID, topicID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoFollowTopic", reflect.TypeOf((*MockTopicRepo)(nil).DoFollowTopic), accountID, topicID)
}

// GetAllTopics mocks base method.
func (m *MockTopicRepo) GetAllTopics() ([]model.Topic, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTopics")
	ret0, _ := ret[0].([]model.Topic)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTopics indicates an expected call of GetAllTopics.
func (mr *MockTopicRepoMockRecorder) GetAllTopics() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTopics", reflect.TypeOf((*MockTopicRepo)(nil).GetAllTopics))
}

// GetFollowingTopicIDs mocks base method.
func (m *MockTopicRepo) GetFollowingTopicIDs(accountID int) ([]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFollowingTopicIDs", accountID)
	ret0, _ := ret[0].([]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFollowingTopicIDs indicates an expected call of GetFollowingTopicIDs.
func (mr *MockTopicRepoMockRecorder) GetFollowingTopicIDs(accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFollowingTopicIDs", reflect.TypeOf((*MockTopicRepo)(nil).GetFollowingTopicIDs), accountID)
}

// GetFollowingTopics mocks base method.
func (m *MockTopicRepo) GetFollowingTopics(accountId int) ([]model.Topic, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFollowingTopics", accountId)
	ret0, _ := ret[0].([]model.Topic)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFollowingTopics indicates an expected call of GetFollowingTopics.
func (mr *MockTopicRepoMockRecorder) GetFollowingTopics(accountId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFollowingTopics", reflect.TypeOf((*MockTopicRepo)(nil).GetFollowingTopics), accountId)
}

// GetPopularTopics mocks base method.
func (m *MockTopicRepo) GetPopularTopics(limit int) ([]model.Topic, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPopularTopics", limit)
	ret0, _ := ret[0].([]model.Topic)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPopularTopics indicates an expected call of GetPopularTopics.
func (mr *MockTopicRepoMockRecorder) GetPopularTopics(limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPopularTopics", reflect.TypeOf((*MockTopicRepo)(nil).GetPopularTopics), limit)
}

// GetRecentPopularArticleIDs mocks base method.
func (m *MockTopicRepo) GetRecentPopularArticleIDs(duration time.Duration, limit int) ([]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecentPopularArticleIDs", duration, limit)
	ret0, _ := ret[0].([]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecentPopularArticleIDs indicates an expected call of GetRecentPopularArticleIDs.
func (mr *MockTopicRepoMockRecorder) GetRecentPopularArticleIDs(duration, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecentPopularArticleIDs", reflect.TypeOf((*MockTopicRepo)(nil).GetRecentPopularArticleIDs), duration, limit)
}

// GetTopicData mocks base method.
func (m *MockTopicRepo) GetTopicData(topicID int) (*model.Topic, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTopicData", topicID)
	ret0, _ := ret[0].(*model.Topic)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTopicData indicates an expected call of GetTopicData.
func (mr *MockTopicRepoMockRecorder) GetTopicData(topicID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTopicData", reflect.TypeOf((*MockTopicRepo)(nil).GetTopicData), topicID)
}

// GetTopicToArticleArrayByArticleID mocks base method.
func (m *MockTopicRepo) GetTopicToArticleArrayByArticleID(articleID int) ([]model.TopicsToArticles, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTopicToArticleArrayByArticleID", articleID)
	ret0, _ := ret[0].([]model.TopicsToArticles)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTopicToArticleArrayByArticleID indicates an expected call of GetTopicToArticleArrayByArticleID.
func (mr *MockTopicRepoMockRecorder) GetTopicToArticleArrayByArticleID(articleID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTopicToArticleArrayByArticleID", reflect.TypeOf((*MockTopicRepo)(nil).GetTopicToArticleArrayByArticleID), articleID)
}

// GetTopicToArticleArrayByArticleIDs mocks base method.
func (m *MockTopicRepo) GetTopicToArticleArrayByArticleIDs(articleIDs []int) ([]model.TopicsToArticles, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTopicToArticleArrayByArticleIDs", articleIDs)
	ret0, _ := ret[0].([]model.TopicsToArticles)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTopicToArticleArrayByArticleIDs indicates an expected call of GetTopicToArticleArrayByArticleIDs.
func (mr *MockTopicRepoMockRecorder) GetTopicToArticleArrayByArticleIDs(articleIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTopicToArticleArrayByArticleIDs", reflect.TypeOf((*MockTopicRepo)(nil).GetTopicToArticleArrayByArticleIDs), articleIDs)
}

// GetTopicsByCollectionID mocks base method.
func (m *MockTopicRepo) GetTopicsByCollectionID(collectionID int) ([]model.Topic, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTopicsByCollectionID", collectionID)
	ret0, _ := ret[0].([]model.Topic)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTopicsByCollectionID indicates an expected call of GetTopicsByCollectionID.
func (mr *MockTopicRepoMockRecorder) GetTopicsByCollectionID(collectionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTopicsByCollectionID", reflect.TypeOf((*MockTopicRepo)(nil).GetTopicsByCollectionID), collectionID)
}

// IsFollowingTopic mocks base method.
func (m *MockTopicRepo) IsFollowingTopic(accountID, topicID int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsFollowingTopic", accountID, topicID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsFollowingTopic indicates an expected call of IsFollowingTopic.
func (mr *MockTopicRepoMockRecorder) IsFollowingTopic(accountID, topicID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsFollowingTopic", reflect.TypeOf((*MockTopicRepo)(nil).IsFollowingTopic), accountID, topicID)
}

// UnfollowTopic mocks base method.
func (m *MockTopicRepo) UnfollowTopic(accountID, topicID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnfollowTopic", accountID, topicID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnfollowTopic indicates an expected call of UnfollowTopic.
func (mr *MockTopicRepoMockRecorder) UnfollowTopic(accountID, topicID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnfollowTopic", reflect.TypeOf((*MockTopicRepo)(nil).UnfollowTopic), accountID, topicID)
}
