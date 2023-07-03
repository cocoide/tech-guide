package repository

import (
	"github.com/cocoide/tech-guide/pkg/model"
	"gorm.io/gorm"
)

type ArticleRepo interface {
	Create(article *model.Article) (int, error)
	CheckArticleExistsByURL(url string) (bool, error)
	BatchCreate(articles []*model.Article) ([]int, error)
	CreateTopicToArticle(topicToArticles []model.TopicsToArticles) error
	GetLatestArticleByLimit(limit int) ([]*model.Article, error)
	GetArticlesByIDs(articleIDs []int) ([]model.Article, error)
	GetArticleByID(articleID int) (*model.Article, error)
	GetTopicsByID(articleId int) ([]model.Topic, error)
	GetTagsAndWeightsByArticleID(articleID int) ([]model.TopicsToArticles, error)
	// Limit by 50 counts
	GetArticlesByTopicIDs(topicIDs []int, omitArticleId int) ([]model.TopicsToArticles, error)
}

type articleRepo struct {
	db *gorm.DB
}

func NewArticleRepo(db *gorm.DB) ArticleRepo {
	return &articleRepo{db: db}
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

func (r *articleRepo) GetLatestArticleByLimit(limit int) ([]*model.Article, error) {
	var articles []*model.Article
	if err := r.db.
		Preload("Topics").
		Order("created_at desc").
		Limit(limit).
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
