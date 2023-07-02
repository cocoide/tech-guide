package service_test

import (
	"reflect"
	"testing"

	"github.com/cocoide/tech-guide/conf"
	"github.com/cocoide/tech-guide/key"
	repoMock "github.com/cocoide/tech-guide/mock/repository"
	"github.com/cocoide/tech-guide/pkg/model"
	"github.com/cocoide/tech-guide/pkg/service"
	"github.com/cocoide/tech-guide/pkg/util"
	"github.com/golang/mock/gomock"
)

func Test_GetRecommendArticleIDs(t *testing.T) {
	conf.NewEnv()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tr := repoMock.NewMockTopicRepo(ctrl)
	cr := repoMock.NewMockCacheRepo(ctrl)
	ps := service.NewPersonalizeService(tr, cr)

	testAccountID := 30
	testTopicIDs := []int{1, 2, 3}
	testPupularArticleIDs := []int{4, 5, 6}
	testTopicToArticleArray := []model.TopicsToArticles{
		{
			ArticleID: 4,
			TopicID:   1,
			Weight:    5,
		},
		// weight 4だけだと0.8になるので選ばれない
		{
			ArticleID: 5,
			TopicID:   2,
			Weight:    4,
		},
		{
			ArticleID: 4,
			TopicID:   2,
			Weight:    5,
		},
		{
			ArticleID: 6,
			TopicID:   3,
			Weight:    4,
		},
		{
			ArticleID: 6,
			TopicID:   3,
			Weight:    3,
		},
	}

	tr.EXPECT().
		GetFollowingTopicIDs(testAccountID).
		Return(testTopicIDs, nil)

	testStrIDs, _ := util.Serialize(testPupularArticleIDs)
	cr.EXPECT().
		Get(key.PopularArticleIDs).
		Return(testStrIDs, nil)

	tr.EXPECT().
		GetTopicToArticleArrayByArticleIDs(testPupularArticleIDs).
		Return(testTopicToArticleArray, nil)

	result, err := ps.GetRecommendArticleIDs(testAccountID)
	if err != nil {
		t.Error(err)
	}
	expected := []int{4, 6}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("exp %v, but got %v", expected, result)
	}
}
