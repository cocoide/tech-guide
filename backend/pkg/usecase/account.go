package usecase

import (
	"errors"

	"github.com/cocoide/tech-guide/pkg/domain/model"
	"github.com/cocoide/tech-guide/pkg/domain/repository"
	"github.com/cocoide/tech-guide/pkg/utils"
)

type AccountUsecase struct {
	repo repository.Repository
}

func NewAccountUsecase(repo repository.Repository) *AccountUsecase {
	return &AccountUsecase{repo: repo}
}

func (au *AccountUsecase) Login(email string) (*model.Account, string, error) {
	account, err := au.repo.GetByEmail(email)
	if err != nil {
		return nil, "", err
	}
	token, err := utils.GenerateToken(account.ID)
	if err != nil {
		return nil, "", err
	}
	return account, token, nil
}

func (au *AccountUsecase) SignUp(account *model.Account) (*model.Account, error) {
	isEmailUsed, err := au.repo.CheckExistByEmail(account.Email)
	if err != nil {
		return nil, err
	}
	if isEmailUsed {
		return nil, errors.New("email is already used")
	}
	account, err = au.repo.CreateAccount(account)
	if err != nil {
		return nil, err
	}
	return account, nil
}
