package rdb

import (
	"errors"

	"github.com/cocoide/tech-guide/pkg/domain/model"
	"gorm.io/gorm"
)

func (r *Repository) CreateFollowTopics(accountID int, topicIDs []int) error {
	var followTopics []*model.FollowTopic
	for _, topicID := range topicIDs {
		followTopics = append(followTopics, &model.FollowTopic{TopicID: topicID, AccountID: accountID})
	}
	return r.db.Create(followTopics).Error
}

func (r *Repository) CreateFollowSources(accountID int, sourceIDs []int) error {
	var followSources []*model.FollowSource
	for _, sourceID := range sourceIDs {
		followSources = append(followSources, &model.FollowSource{SourceID: sourceID, AccountID: accountID})
	}
	return r.db.Create(followSources).Error
}

func (r *Repository) GetAllAccountIDs() ([]int, error) {
	var accountIDs []int
	if err := r.db.Model(&model.Account{}).Pluck("id", &accountIDs).Error; err != nil {
		return nil, err
	}
	return accountIDs, nil
}

func (r *Repository) GetAccountProfile(id int) (*model.Account, error) {
	var account model.Account
	err := r.db.
		Where("id=?", id).
		First(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *Repository) CreateAccount(account *model.Account) (*model.Account, error) {
	if err := r.db.Create(account).Error; err != nil {
		return nil, err
	}
	return account, nil
}

func (r *Repository) CheckExistByEmail(email string) (bool, error) {
	var account model.Account
	if err := r.db.
		Where("email=?", email).
		First(&account).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
func (r *Repository) GetByEmail(email string) (*model.Account, error) {
	var account model.Account
	if err := r.db.
		Where("email=?", email).
		First(&account).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &account, nil
}
