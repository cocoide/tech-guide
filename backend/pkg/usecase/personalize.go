package usecase

import (
	"time"

	"github.com/cocoide/tech-guide/key"
	"github.com/cocoide/tech-guide/pkg/domain/repository"
	"github.com/cocoide/tech-guide/pkg/utils"
)

type PersonalizeUsecase struct {
	repo  repository.Repository
	cache repository.CacheRepo
}

func NewPersonalizeUsecase(repo repository.Repository, cache repository.CacheRepo) *PersonalizeUsecase {
	return &PersonalizeUsecase{repo: repo, cache: cache}
}

func (s *PersonalizeUsecase) GetRecommendArticleIDs(accoutId int) ([]int, error) {
	topicIDs, err := s.repo.GetFollowingTopicIDs(accoutId)
	if err != nil {
		return nil, err
	}
	if len(topicIDs) < 1 {
		return nil, nil
	}
	// パフォーマンス向上のためにTopicIDsのMapを作成
	topicIDMap := make(map[int]struct{})
	for _, topicID := range topicIDs {
		topicIDMap[topicID] = struct{}{}
	}
	var articleIDs []int
	strArticleIDs, err := s.cache.Get(key.PopularArticleIDs)
	if len(strArticleIDs) < 1 || err != nil {
		articleIDs, err = s.repo.GetRecentPopularArticleIDs(30*24*time.Hour, 30)
		serializedIDs, err := utils.Serialize(articleIDs)
		if err != nil {
			return nil, err
		}
		if err := s.cache.Set(key.PopularArticleIDs, serializedIDs, 30*24*time.Hour); err != nil {
			return nil, err
		}
	} else {
		articleIDs, err = utils.Deserialize[[]int](strArticleIDs)
	}
	topicToArticles, err := s.repo.GetTopicToArticleArrayByArticleIDs(articleIDs)
	// フォロー中のトピックに基づいて記事の推奨度を計算
	recommends := make(map[int]float64, len(articleIDs)) // articleID -> relevance
	for _, v := range topicToArticles {
		if _, ok := topicIDMap[v.TopicID]; ok {
			_, ok := recommends[v.ArticleID]
			if v.Weight <= 1 {
				continue
			}
			normalizedWeight := float64(v.Weight) / 5
			if ok {
				if recommends[v.ArticleID] < normalizedWeight {
					recommends[v.ArticleID] = normalizedWeight + recommends[v.ArticleID]/10
				} else {
					recommends[v.ArticleID] += normalizedWeight / 10
				}
			} else {
				recommends[v.ArticleID] = normalizedWeight
			}
		}
	}
	// 降順に並び替え、閾値を0.8より大きい値で設定
	result := utils.SortMapKeysByFloatValue[int](recommends, 0.8)
	return result, nil
}
