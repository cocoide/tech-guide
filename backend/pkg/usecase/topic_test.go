package usecase

import (
	"github.com/cocoide/tech-guide/conf"
	"github.com/cocoide/tech-guide/pkg/infrastructure/database"
	"github.com/cocoide/tech-guide/pkg/infrastructure/integration"
	"github.com/cocoide/tech-guide/pkg/infrastructure/rdb"
	"testing"
)

func Test_TopicAssign(t *testing.T) {
	//ctrl := gomock.NewController(t)
	//defer ctrl.Finish()
	//articleRepo := mock_repository.NewMockArticleRepo(ctrl)
	//topicRepo := mock_repository.NewMockTopicRepo(ctrl)
	//nlpServ := mock_service.NewMockNLPService(ctrl)
	conf.NewEnv()
	gorm := database.NewGormClient()
	repo := rdb.NewRepository(gorm)
	nlp := integration.NewNLPService()
	scraper := integration.NewScrapingService()

	//type TopicUsecaseRepo struct {
	//	repository.ArticleRepo
	//	repository.TopicRepo
	//}
	usecase := NewTopicUsecase(repo, nlp, scraper)
	usecase.UpsertTopicToArticlesByArticleID(1)
}
