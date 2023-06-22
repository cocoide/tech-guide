package usecase

import (
	"errors"

	"github.com/cocoide/tech-guide/pkg/model"
	repo "github.com/cocoide/tech-guide/pkg/repository"
	"github.com/cocoide/tech-guide/pkg/util"
)

type AccountUseCase interface {
	Login(email string) (*model.Account, string, error)
	SignUp(account *model.Account) (*model.Account, error)
}
type accountUseCase struct {
	ur repo.AccountRepo
}

func NewAccountUseCase(ur repo.AccountRepo) AccountUseCase {
	return &accountUseCase{ur: ur}
}

func (au *accountUseCase) Login(email string) (*model.Account, string, error) {
	account, err := au.ur.GetByEmail(email)
	if err != nil {
		return nil, "", err
	}
	token, err := util.GenerateToken(account.ID)
	if err != nil {
		return nil, "", err
	}
	return account, token, nil
}

func (au *accountUseCase) SignUp(account *model.Account) (*model.Account, error) {
	isEmailUsed, err := au.ur.CheckExistByEmail(account.Email)
	if err != nil {
		return nil, err
	}
	if isEmailUsed {
		return nil, errors.New("email is already used")
	}
	if len(account.Password) > 0 {
		hashed, err := util.HashPassword(account.Password)
		if err != nil {
			return nil, err
		}
		account.Password = hashed
	}
	account, err = au.ur.Create(account)
	if err != nil {
		return nil, err
	}
	return account, nil
}
