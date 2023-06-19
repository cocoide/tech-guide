package usecase

import (
	"errors"

	"github.com/cocoide/tech-guide/pkg/model"
	repo "github.com/cocoide/tech-guide/pkg/repository"
	"github.com/cocoide/tech-guide/pkg/util"
)

type AccountUseCase interface {
	Login(email, password string) (*model.Account, string, error)
	SignUp(email, password string) error
}
type accountUseCase struct {
	ur repo.AccountRepo
}

func NewAccountUseCase(ur repo.AccountRepo) AccountUseCase {
	return &accountUseCase{ur: ur}
}

func (au *accountUseCase) Login(email, password string) (*model.Account, string, error) {
	account, err := au.ur.GetByEmail(email)
	if err != nil {
		return nil, "", err
	}
	if err := util.CheckPassword(account.Password, password); err != nil {
		return nil, "", err
	}
	token, err := util.GenerateToken(account.ID)
	if err != nil {
		return nil, "", err
	}
	return account, token, nil
}

func (au *accountUseCase) SignUp(email, password string) error {
	isEmailUsed, err := au.ur.CheckExistByEmail(email)
	if err != nil {
		return err
	}
	if isEmailUsed {
		return errors.New("email is already used")
	}
	hashed, err := util.HashPassword(password)
	if err != nil {
		return err
	}
	account := &model.Account{
		Email:    email,
		Password: hashed,
	}
	if err := au.ur.Create(account); err != nil {
		return err

	}
	return nil
}
