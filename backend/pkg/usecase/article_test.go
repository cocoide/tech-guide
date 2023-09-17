package usecase_test

import (
	"context"
	"testing"

	"github.com/cocoide/tech-guide/conf"
	mock_repository "github.com/cocoide/tech-guide/mock/repository"
	"github.com/cocoide/tech-guide/pkg/domain/model"
	"github.com/cocoide/tech-guide/pkg/infrastructure/integration"
	"github.com/cocoide/tech-guide/pkg/usecase"

	"github.com/golang/mock/gomock"
)

func Test_ExtractTopicIDFromArticleTitle(t *testing.T) {
	conf.NewEnv()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tr := mock_repository.NewMockTopicRepo(ctrl)
	topics := []model.Topic{
		{ID: 1, Name: "個人開発"},
		{ID: 2, Name: "Marketing"},
		{ID: 3, Name: "Golang"},
		{ID: 4, Name: "Ruby"},
		{ID: 5, Name: "Businness"},
		{ID: 6, Name: "設計"},
		{ID: 7, Name: "CSS"},
		{ID: 8, Name: "フロントエンド"},
		{ID: 9, Name: "可読性"},
	}
	tr.EXPECT().
		GetAllTopics().
		Return(topics, nil)

	ctx := context.Background()
	openai :=integration.NewOpenAIService(ctx)
	ts := usecase.NewArticleUsecase(openai,tr)
	topicWeights, err := ts.ExtractTopicsWithWeightFromArticleTitle("Tailwind CSSはCSS設計に何をもたらすか")
	if err != nil {
		t.Error(err)
	}
	t.Logf("result is :%v", topicWeights)
}
