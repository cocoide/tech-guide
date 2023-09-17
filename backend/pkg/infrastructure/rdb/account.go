package rdb

import (
	"errors"

	"github.com/cocoide/tech-guide/pkg/domain/model"
	"gorm.io/gorm"
)

//type Repository struct {
//	db *gorm.DB
//}
//
//func NewRepository(db *gorm.DB) repository.AccountRepo {
//	return &Repository{db: db}
//}

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
