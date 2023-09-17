package rdb

import (
	"github.com/cocoide/tech-guide/pkg/domain/model"
)


func (r *Repository) FindContributionsByAccountID(accountID int) ([]*model.Contribution, error) {
	var contribution []*model.Contribution
	if err := r.db.Where("account_id = ?", accountID).Find(&contribution).Error; err != nil {
		return nil, err
	}
	return contribution, nil
}

func (r *Repository) UpdateContribution(contribute *model.Contribution) error {
	return r.db.Updates(contribute).Error
}

func (r *Repository) CreateContribution(contribute *model.Contribution) error {
	return r.db.Create(contribute).Error
}

func (r *Repository) CreateAchievement(achieve *model.AccountsToAchievements) error {
	return r.db.Create(achieve).Error
}

func (r *Repository) BatchCreateContributions(contributes []*model.Contribution) error {
	return r.db.Create(contributes).Error
}

func (r *Repository) GetContributionsByAccountID(accountID int) ([]*model.Contribution, error) {
	var contributions []*model.Contribution
	if err := r.db.Where("account_id = ?", accountID).Find(&contributions).Error; err != nil {
		return nil, err
	}
	return contributions, nil
}
