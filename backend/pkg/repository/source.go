package repository

import (
	"github.com/cocoide/tech-guide/pkg/model"
	"gorm.io/gorm"
)

type SourceRepo interface {
	// 存在しない場合は0を返す
	FindIDByDomain(domain string) (int, error)
	FollowSource(followSource *model.FollowSource) error
	GetPopularSources(limit int) ([]model.Source, error)
}

type sourceRepo struct {
	db *gorm.DB
}

func NewSourceRepo(db *gorm.DB) SourceRepo {
	return &sourceRepo{db: db}
}

func (r *sourceRepo) GetPopularSources(limit int) ([]model.Source, error) {
	var popularSources []model.Source
	err := r.db.
		Table("follow_sources").
		Select("sources.*, COUNT(follow_sources.id) as follow_count").
		Joins("INNER JOIN sources ON sources.id = follow_sources.source_id").
		Group("sources.id").
		Order("follow_count DESC").
		Limit(limit).
		Find(&popularSources).
		Error
	if err != nil {
		return nil, err
	}

	return popularSources, nil
}

func (r *sourceRepo) FindIDByDomain(domain string) (int, error) {
	var sourceID int
	err := r.db.
		Table("sources").
		Where("domain = ?", domain).
		Pluck("id", &sourceID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, nil
		}
		return 0, err
	}
	return sourceID, nil
}

func (r *sourceRepo) FollowSource(followSource *model.FollowSource) error {
	return r.db.Create(followSource).Error
}
