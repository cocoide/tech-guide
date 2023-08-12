package repository

import (
	"github.com/cocoide/tech-guide/pkg/model"
	"gorm.io/gorm"
)

type SourceRepo interface {
	// 存在しない場合は0を返す
	FindIDByDomain(domain string) (int, error)
	GetPopularSources(limit int) ([]model.Source, error)
	GetAllSources() ([]model.Source, error)
	UnFollowSource(accountID, sourceID int) error
	DoFollowSource(accountID, souceID int) error
	GetFollowingSources(accountId int) ([]model.Source, error)
	GetFollowingSourceIDs(accountID int) ([]int, error)
	GetSourceData(sourceID int) (*model.Source, error)
	IsFollowingSource(accountID, sourceID int) (bool, error)
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

func (r *sourceRepo) GetAllSources() ([]model.Source, error) {
	var souces []model.Source
	if err := r.db.Find(&souces).Error; err != nil {
		return nil, err
	}
	return souces, nil
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

func (r *sourceRepo) UnFollowSource(accountID, sourceID int) error {
	return r.db.Where("account_id = ? AND source_id = ?", accountID, sourceID).
		Delete(&model.FollowSource{}).Error
}

func (r *sourceRepo) DoFollowSource(accountID, souceID int) error {
	followSource := model.FollowSource{
		AccountID: accountID,
		SourceID:  souceID,
	}
	return r.db.Create(&followSource).Error
}
func (r *sourceRepo) IsFollowingSource(accountID, sourceID int) (bool, error) {
	var count int64
	if err := r.db.Model(&model.FollowSource{}).
		Where("account_id = ? AND source_id = ?", accountID, sourceID).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (tr *sourceRepo) GetFollowingSources(accountId int) ([]model.Source, error) {
	account := &model.Account{}
	err := tr.db.Preload("FollowSources").
		First(account, accountId).Error
	if err != nil {
		return nil, err
	}
	return account.FollowSources, nil
}

func (tr *sourceRepo) GetFollowingSourceIDs(accountID int) ([]int, error) {
	var sourceIDs []int
	err := tr.db.Model(&model.FollowSource{}).
		Where("account_id = ?", accountID).
		Pluck("source_id", &sourceIDs).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return sourceIDs, nil
}

func (r *sourceRepo) GetSourceData(sourceID int) (*model.Source, error) {
	var source model.Source
	if err := r.db.First(&source, sourceID).Error; err != nil {
		return nil, err
	}
	return &source, nil
}
