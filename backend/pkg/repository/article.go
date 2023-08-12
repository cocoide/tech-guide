package repository

import (
	"errors"

	"github.com/cocoide/tech-guide/pkg/model"
	"gorm.io/gorm"
)

type ArticleRepo interface {
	Create(article *model.Article) (int, error)
	GetArticleIDByURL(url string) (int, error)
	CheckArticleExistsByURL(url string) (bool, error)
	BatchCreate(articles []*model.Article) ([]int, error)
	CreateTopicToArticle(topicToArticles []model.TopicsToArticles) error
	// summaryカラムはomit
	GetLatestArticleByLimitWithSourceData(pageIndex, pageSize int) ([]*model.Article, error)
	GetArticlesByIDs(articleIDs []int) ([]model.Article, error)
	GetArticleByID(articleID int) (*model.Article, error)
	GetArticleWithRelatedDataByID(articleID int) (*model.Article, error)
	GetTopicsByID(articleId int) ([]model.Topic, error)
	GetTagsAndWeightsByArticleID(articleID int) ([]model.TopicsToArticles, error)
	// Limit by 50 counts
	GetArticlesByTopicIDs(topicIDs []int, omitArticleId int) ([]model.TopicsToArticles, error)
	UpdateSummaryByID(id int, summary string) error
	BatchGetArticlesByTopicIDsAndSourceID(topicIDs, sourceIDs []int, pageIndex, pageSize int) ([]model.Article, error)
	GetArticlesBySourceID(sourceID, pageIndex, pageSize int) ([]model.Article, error)
	GetArticlesByTopicID(topicID, pageIndex, pageSize int) ([]model.Article, error)
}

type articleRepo struct {
	db *gorm.DB
}

func NewArticleRepo(db *gorm.DB) ArticleRepo {
	return &articleRepo{db: db}
}

func (r *articleRepo) GetArticleIDByURL(url string) (int, error) {
	var id int
	err := r.db.Model(&model.Article{}).Where("original_url = ?", url).Pluck("id", &id).Error
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *articleRepo) UpdateSummaryByID(id int, summary string) error {
	if err := r.db.Model(&model.Article{}).Where("id = ?", id).Update("summary", summary).Error; err != nil {
		return err
	}
	return nil
}

func (r *articleRepo) CheckArticleExistsByURL(url string) (bool, error) {
	var count int64
	err := r.db.Model(&model.Article{}).Where("original_url = ?", url).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *articleRepo) GetArticlesByIDs(articleIDs []int) ([]model.Article, error) {
	var articles []model.Article
	err := r.db.
		Where("id IN (?)", articleIDs).
		Preload("Topics").
		Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (r *articleRepo) GetArticleWithRelatedDataByID(articleID int) (*model.Article, error) {
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

func (r *articleRepo) GetArticleByID(articleID int) (*model.Article, error) {
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
func (r *articleRepo) CreateTopicToArticle(topicToArticles []model.TopicsToArticles) error {
	return r.db.Create(&topicToArticles).Error
}

func (r *articleRepo) Create(article *model.Article) (int, error) {
	if err := r.db.Create(article).Error; err != nil {
		return 0, err
	}
	return article.ID, nil
}
func (r *articleRepo) BatchCreate(articles []*model.Article) ([]int, error) {
	if err := r.db.Create(articles).Error; err != nil {
		return nil, err
	}
	var result []int
	for _, v := range articles {
		result = append(result, v.ID)
	}
	return result, nil
}

func (r *articleRepo) GetLatestArticleByLimitWithSourceData(pageIndex, pageSize int) ([]*model.Article, error) {
	if pageIndex < 1 {
		return nil, errors.New("page number should start from 1")
	}
	var articles []*model.Article
	if err := r.db.
		Omit("summary").
		Preload("Source").
		Order("created_at desc").
		Offset((pageIndex - 1) * pageSize).
		Limit(pageSize).
		Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}
func (r *articleRepo) GetTopicsByID(articleId int) ([]model.Topic, error) {
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

func (r *articleRepo) GetTagsAndWeightsByArticleID(articleID int) ([]model.TopicsToArticles, error) {
	var TopicsToArticles []model.TopicsToArticles
	if err := r.db.
		Preload("Topic").
		Where("article_id IN (?)", articleID).
		Find(&TopicsToArticles).Error; err != nil {
		return nil, err
	}
	return TopicsToArticles, nil
}

func (r *articleRepo) GetArticlesByTopicIDs(topicIDs []int, omitArticleId int) ([]model.TopicsToArticles, error) {
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

func (r *articleRepo) BatchGetArticlesByTopicIDsAndSourceID(topicIDs, sourceIDs []int, pageIndex, pageSize int) ([]model.Article, error) {
	var articles []model.Article
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
func (r *articleRepo) GetArticlesBySourceID(sourceID, pageIndex, pageSize int) ([]model.Article, error) {
	var articles []model.Article

	err := r.db.
		Preload("Source").
		Where("source_id = ?", sourceID).
		Offset((pageIndex - 1) * pageSize).
		Find(&articles).Error
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (r *articleRepo) GetArticlesByTopicID(topicID, pageIndex, pageSize int) ([]model.Article, error) {
	var topicsToArticles []model.TopicsToArticles
	var articles []model.Article

	err := r.db.Preload("Article").
		Where("topic_id = ?", topicID).
		Offset((pageIndex - 1) * pageSize).
		Find(&topicsToArticles).Error
	if err != nil {
		return nil, err
	}
	for _, ta := range topicsToArticles {
		articles = append(articles, ta.Article)
	}
	return articles, nil
}
