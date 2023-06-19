package repository

import (
	"errors"

	"github.com/cocoide/tech-guide/pkg/model"
	"gorm.io/gorm"
)

type AccountRepo interface {
	Create(account *model.Account) error
	CheckExistByEmail(email string) (bool, error)
	GetByEmail(email string) (*model.Account, error)
	GetAccountProfile(id int) (*model.Account, error)
}

type accountRepo struct {
	db *gorm.DB
}

func NewAccountRepo(db *gorm.DB) AccountRepo {
	return &accountRepo{db: db}
}

func (ar *accountRepo) GetAccountProfile(id int) (*model.Account, error) {
	var account model.Account
	err := ar.db.
		Where("id=?", id).
		First(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (ar *accountRepo) Create(account *model.Account) error {
	return ar.db.Create(account).Error
}

func (ar *accountRepo) CheckExistByEmail(email string) (bool, error) {
	var account model.Account
	if err := ar.db.
		Where("email=?", email).
		First(&account).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
func (ar *accountRepo) GetByEmail(email string) (*model.Account, error) {
	var account model.Account
	if err := ar.db.
		Where("email=?", email).
		First(&account).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &account, nil
}
