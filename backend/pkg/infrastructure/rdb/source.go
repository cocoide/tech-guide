package rdb

import (
	"github.com/cocoide/tech-guide/pkg/domain/model"
	"gorm.io/gorm"
)

func (r *Repository) GetPopularSources(limit int) ([]model.Source, error) {
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

func (r *Repository) GetAllSources() ([]model.Source, error) {
	var souces []model.Source
	if err := r.db.Find(&souces).Error; err != nil {
		return nil, err
	}
	return souces, nil
}

func (r *Repository) FindIDByDomain(domain string) (int, error) {
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

func (r *Repository) UnFollowSource(accountID, sourceID int) error {
	return r.db.Where("account_id = ? AND source_id = ?", accountID, sourceID).
		Delete(&model.FollowSource{}).Error
}

func (r *Repository) DoFollowSource(accountID, souceID int) error {
	followSource := model.FollowSource{
		AccountID: accountID,
		SourceID:  souceID,
	}
	return r.db.Create(&followSource).Error
}
func (r *Repository) IsFollowingSource(accountID, sourceID int) (bool, error) {
	var count int64
	if err := r.db.Model(&model.FollowSource{}).
		Where("account_id = ? AND source_id = ?", accountID, sourceID).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (tr *Repository) GetFollowingSources(accountId int) ([]model.Source, error) {
	account := &model.Account{}
	err := tr.db.Preload("FollowSources").
		First(account, accountId).Error
	if err != nil {
		return nil, err
	}
	return account.FollowSources, nil
}

func (tr *Repository) GetFollowingSourceIDs(accountID int) ([]int, error) {
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

func (r *Repository) GetSourceData(sourceID int) (*model.Source, error) {
	var source model.Source
	if err := r.db.First(&source, sourceID).Error; err != nil {
		return nil, err
	}
	return &source, nil
}
