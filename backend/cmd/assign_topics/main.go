package main

import (
	"flag"
	"github.com/cocoide/tech-guide/conf"
	"github.com/cocoide/tech-guide/pkg/domain/model"
	"github.com/cocoide/tech-guide/pkg/infrastructure/database"
	"github.com/cocoide/tech-guide/pkg/infrastructure/integration"
	"github.com/cocoide/tech-guide/pkg/infrastructure/rdb"
	"github.com/cocoide/tech-guide/pkg/usecase"
	"log"
	"strings"
)

const (
	mediumWeight = 3
)

type AssignTopicsType int

const (
	Create AssignTopicsType = iota + 1
	Update
)

func main() {
	conf.NewEnv()
	db := database.NewGormClient()
	repo := rdb.NewRepository(db)
	nlp := integration.NewNLPService()
	scraper := integration.NewScrapingService()
	topicUsecase := usecase.NewTopicUsecase(repo, nlp, scraper)
	excludeIDs := []int{0}
	targetPtr := flag.Int("type", 0, "usecase type")
	flag.Parse()
	target := *targetPtr
	targetType := AssignTopicsType(target)
	if targetType == 0 {
		log.Fatal("Require to select -type")
	}

	switch targetType {

	case Update:
		articles, _ := repo.GetArticlesExcludingIDs(excludeIDs)
		topics, _ := repo.GetAllTopics()
		var topicsMap map[string]value

		for _, article := range articles {
			topicsMap = CountTopicsInSentence(article.Title+article.Summary, topics)
			var topicToArticles []model.TopicsToArticles
			found := false
			for _, v := range topicsMap {
				if topicsMap != nil {
					topicToArticles = append(topicToArticles,
						model.TopicsToArticles{
							TopicID:   v.id,
							Weight:    mediumWeight,
							ArticleID: article.ID,
						})
					found = true
				}
			}
			if found {
				if err := repo.CreateTopicToArticle(topicToArticles); err != nil {
					log.Print(err)
				}
			}
		}

	case Create:
		articles, _ := repo.GetArticlesExcludingIDs([]int{1})
		var articleIDs []int
		for _, v := range articles[0:50] {
			articleIDs = append(articleIDs, v.ID)
		}

		articleIDTopicMap := make(map[int][]int, 0)
		articleTopics, _ := repo.GetTopicToArticleArrayByArticleIDs(articleIDs)
		for _, articleTopic := range articleTopics {
			if topicIDs, ok := articleIDTopicMap[articleTopic.ArticleID]; ok {
				topicIDs = append(topicIDs, articleTopic.TopicID)
				articleIDTopicMap[articleTopic.ArticleID] = topicIDs
			} else {
				articleIDTopicMap[articleTopic.ArticleID] = []int{articleTopic.TopicID}
			}
		}
		log.Print(articleIDTopicMap)
		categories, _ := repo.GetCategoriesWithTopics()
		categoryNames := make([]string, len(categories))
		for _, c := range categories {
			categoryNames = append(categoryNames, c.Name)
		}
		topicUsecase.UpsertTopicToArticlesByArticleID(235)

	}

}

type value struct {
	id    int
	count int
}

func normalize(s string) string {
	return strings.ToLower(s) // すべて小文字に変換
}

func CountTopicsInSentence(sentence string, topics []model.Topic) map[string]value {
	normSentence := normalize(sentence)
	topicsMap := make(map[string]value, len(topics))
	for _, topic := range topics {
		normTopicName := normalize(topic.Name)
		count := strings.Count(normSentence, normTopicName)
		if count > 0 {
			topicsMap[normTopicName] = value{id: topic.ID, count: count}
		}
	}
	return topicsMap
}
