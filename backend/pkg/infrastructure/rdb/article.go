package rdb

import (
	"errors"
	"fmt"
	repo "github.com/cocoide/tech-guide/pkg/domain/repository"
	"log"

	"github.com/cocoide/tech-guide/pkg/domain/model"
)

func (r *Repository) ListArticleIDs(params *repo.ListArticlesParams) ([]int, error) {
	var articleIDs []int
	query := r.db.Model(&model.Article{}).Select("articles.id")

	if len(params.SourceIDs) > 0 {
		query = query.Where("source_id IN (?)", params.SourceIDs)
	}
	if len(params.TopicIDs) > 0 {
		query = query.Joins("JOIN topics_to_articles ON articles.id = topics_to_articles.article_id").
			Where("topics_to_articles.topic_id IN (?)", params.TopicIDs)
	}
	limit := 50 // Default
	if params.Limit != 0 {
		limit = params.Limit
	}
	query = query.Limit(limit)
	switch params.OrderBy {
	case repo.Latest:
		query = query.Order("articles.created_at desc")
	case repo.Older:
		query = query.Order("articles.created_at asc")
	default:
		// Default: NoOrder
	}
	log.Print(query)
	if err := query.Pluck("articles.id", &articleIDs).Error; err != nil {
		return nil, err
	}
	log.Print(articleIDs)
	return articleIDs, nil
}

func (r *Repository) ListArticles(params *repo.ListArticlesParams) (model.Articles, error) {
	var articles model.Articles
	query := r.db.Model(&model.Article{})

	for _, v := range params.Preloads {
		query = query.Preload(v)
	}
	if len(params.SourceIDs) > 0 {
		query = query.Where("source_id IN (?)", params.SourceIDs)
	}
	if len(params.TopicIDs) > 0 {
		query = query.Joins("JOIN topics_to_articles ON articles.id = topics_to_articles.article_id").
			Where("topics_to_articles.topic_id IN (?)", params.TopicIDs)
	}
	limit := 50 // Default
	if params.Limit != 0 {
		limit = params.Limit
	}
	query = query.Limit(limit)
	switch params.OrderBy {
	case repo.Latest:
		query = query.Order("created_at desc")
	case repo.Older:
		query = query.Order("created_at asc")
	default:
		// Default: NoOrder
	}
	if err := query.Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}

func (r *Repository) GetArticleIDByURL(url string) (int, error) {
	var id int
	err := r.db.Model(&model.Article{}).Where("original_url = ?", url).Pluck("id", &id).Error
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Repository) UpdateSummaryByID(id int, summary string) error {
	if err := r.db.Model(&model.Article{}).Where("id = ?", id).Update("summary", summary).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) CheckArticleExistsByURL(url string) (bool, error) {
	var count int64
	err := r.db.Model(&model.Article{}).Where("original_url = ?", url).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *Repository) GetArticlesByIDs(articleIDs []int, preloadColumns []string) (model.Articles, error) {
	var articles model.Articles

	query := r.db.Where("id IN (?)", articleIDs)

	for _, column := range preloadColumns {
		query = query.Preload(column)
	}

	err := query.Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (r *Repository) GetArticleWithRelatedDataByID(articleID int) (*model.Article, error) {
	var article model.Article
	err := r.db.
		Where("id = ?", articleID).
		Preload("Topics").
		Preload("Source").
		First(&article).Error
	if err != nil {
		return nil, err
	}
	return &article, nil
}

func (r *Repository) GetArticleByID(articleID int) (*model.Article, error) {
	var article model.Article
	err := r.db.
		Where("id = ?", articleID).
		Preload("Topics").
		First(&article).Error
	if err != nil {
		return nil, err
	}
	return &article, nil
}
func (r *Repository) CreateTopicToArticle(topicToArticles []model.TopicsToArticles) error {
	return r.db.Create(&topicToArticles).Error
}

func (r *Repository) CreateArticle(article *model.Article) (int, error) {
	if err := r.db.Create(article).Error; err != nil {
		return 0, err
	}
	return article.ID, nil
}
func (r *Repository) BatchCreate(articles []*model.Article) ([]int, error) {
	if err := r.db.Create(articles).Error; err != nil {
		return nil, err
	}
	var result []int
	for _, v := range articles {
		result = append(result, v.ID)
	}
	return result, nil
}

func (r *Repository) GetLatestArticleByLimitWithSourceData(pageIndex, pageSize int) ([]*model.Article, error) {
	if pageIndex < 1 {
		return nil, errors.New("page number should start from 1")
	}
	var articles []*model.Article
	if err := r.db.
		//Omit("summary").
		Preload("Source").
		Preload("Rating").
		Order("created_at desc").
		Offset((pageIndex - 1) * pageSize).
		Limit(pageSize).
		Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}
func (r *Repository) GetTopicsByID(articleId int) ([]model.Topic, error) {
	var article model.Article
	if err := r.db.
		Preload("Topics").
		Where("id = ?", articleId).
		First(&article).Error; err != nil {
		return nil, err
	}
	topics := article.Topics
	return topics, nil
}

func (r *Repository) GetTagsAndWeightsByArticleID(articleID int) ([]model.TopicsToArticles, error) {
	var TopicsToArticles []model.TopicsToArticles
	if err := r.db.
		Preload("Topic").
		Where("article_id IN (?)", articleID).
		Find(&TopicsToArticles).Error; err != nil {
		return nil, err
	}
	return TopicsToArticles, nil
}

func (r *Repository) GetArticlesByTopicIDs(topicIDs []int, omitArticleId int) ([]model.TopicsToArticles, error) {
	var TopicsToArticlesArray []model.TopicsToArticles
	if err := r.db.
		Preload("Article").
		Where("topic_id IN (?)", topicIDs).
		Not("article_id = ?", omitArticleId).
		Order("weight DESC").
		Limit(50).
		Find(&TopicsToArticlesArray).Error; err != nil {
		return nil, err
	}
	return TopicsToArticlesArray, nil
}

func (r *Repository) BatchGetArticlesByTopicIDsAndSourceID(topicIDs, sourceIDs []int, pageIndex, pageSize int) (model.Articles, error) {
	var articles model.Articles
	query := r.db.Preload("Topics").Preload("Source").Where("source_id IN (?)", sourceIDs)
	if len(topicIDs) > 0 {
		query = query.Joins("JOIN topics_to_articles ON articles.id = topics_to_articles.article_id").
			Where("topics_to_articles.topic_id IN (?)", topicIDs)
	}
	offset := (pageIndex - 1) * pageSize
	query = query.Offset(offset).Limit(pageSize)
	if err := query.Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}
func (r *Repository) GetArticlesBySourceID(sourceID, pageIndex, pageSize int) ([]model.Article, error) {
	var articles []model.Article
	err := r.db.
		Preload("Rating").
		Preload("Source").
		Where("source_id = ?", sourceID).
		Offset((pageIndex - 1) * pageSize).
		Order("created_at DESC").
		Find(&articles).Error
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (r *Repository) GetArticlesByTopicID(topicID, pageIndex, pageSize int) ([]model.Article, error) {
	var articles []model.Article

	offset := (pageIndex - 1) * pageSize

	err := r.db.
		Joins("JOIN topics_to_articles ON topics_to_articles.article_id = articles.id").
		Preload("Rating").
		Preload("Source").
		Where("topics_to_articles.topic_id = ?", topicID).
		// 関連順に取得する
		Order("topics_to_articles.weight DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&articles).Error
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (r *Repository) FindArticlesByTitle(title string) (articles []model.Article, err error) {
	likeQuery := fmt.Sprintf("%%%s%%", title)
	if err = r.db.
		Where("title LIKE ?", likeQuery).
		Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}

func (r *Repository) GetArticlesExcludingIDs(excludeIDs []int) (model.Articles, error) {
	var articles model.Articles
	err := r.db.Not("id", excludeIDs).Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}
