package repository

import (
	"github.com/cocoide/tech-guide/pkg/model"
	"gorm.io/gorm"
)

type ActivityRepo interface {
	BatchCreateContributions(contribute []*model.Contribution) error
	CreateContribution(contribute *model.Contribution) error
	UpdateContribution(contribute *model.Contribution) error
	CreateAchievement(achieve *model.AccountsToAchievements) error
	GetContributionsByAccountID(accountID int) ([]*model.Contribution, error)
}
type activityRepo struct {
	db *gorm.DB
}

func NewActivityRepo(db *gorm.DB) ActivityRepo {
	return &activityRepo{db: db}
}

func (r *activityRepo) FindContributionsByAccountID(accountID int) ([]*model.Contribution, error) {
	var contribution []*model.Contribution
	if err := r.db.Where("account_id = ?", accountID).Find(&contribution).Error; err != nil {
		return nil, err
	}
	return contribution, nil
}

func (r *activityRepo) UpdateContribution(contribute *model.Contribution) error {
	return r.db.Updates(contribute).Error
}

func (r *activityRepo) CreateContribution(contribute *model.Contribution) error {
	return r.db.Create(contribute).Error
}

func (r *activityRepo) CreateAchievement(achieve *model.AccountsToAchievements) error {
	return r.db.Create(achieve).Error
}

func (r *activityRepo) BatchCreateContributions(contributes []*model.Contribution) error {
	return r.db.Create(contributes).Error
}

func (r *activityRepo) GetContributionsByAccountID(accountID int) ([]*model.Contribution, error) {
	var contributions []*model.Contribution
	if err := r.db.Where("account_id = ?", accountID).Find(&contributions).Error; err != nil {
		return nil, err
	}
	return contributions, nil
}
