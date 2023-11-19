package repository

import (
	"github.com/cocoide/tech-guide/pkg/domain/model"
)

//go:generate mockgen -source=account.go -destination=../../../mock/repository/account.go
type AccountRepo interface {
	CreateAccount(account *model.Account) (*model.Account, error)
	GetAllAccountIDs() ([]int, error)
	CheckExistByEmail(email string) (bool, error)
	GetByEmail(email string) (*model.Account, error)
	GetAccountProfile(id int) (*model.Account, error)
	CreateFollowTopics(accountID int, topicIDs []int) error
	CreateFollowSources(accountID int, sourceIDs []int) error
}
