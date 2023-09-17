package repository

import (
	"github.com/cocoide/tech-guide/pkg/domain/model"
)

//go:generate mockgen -source=activity.go -destination=../../../mock/repository/activity.go
type ActivityRepo interface {
	BatchCreateContributions(contribute []*model.Contribution) error
	CreateContribution(contribute *model.Contribution) error
	UpdateContribution(contribute *model.Contribution) error
	CreateAchievement(achieve *model.AccountsToAchievements) error
	GetContributionsByAccountID(accountID int) ([]*model.Contribution, error)
}