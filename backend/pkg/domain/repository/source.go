package repository

import "github.com/cocoide/tech-guide/pkg/domain/model"

//go:generate mockgen -source=source.go -destination=../../../mock/repository/source.go
type SourceRepo interface {
	// 存在しない場合は0を返す
	FindIDByDomain(domain string) (int, error)
	GetPopularSources(limit int) ([]model.Source, error)
	GetAllSources() ([]model.Source, error)
	UnFollowSource(accountID, sourceID int) error
	DoFollowSource(accountID, souceID int) error
	GetFollowingSources(accountId int) ([]model.Source, error)
	GetFollowingSourceIDs(accountID int) ([]int, error)
	GetSourceData(sourceID int) (*model.Source, error)
	IsFollowingSource(accountID, sourceID int) (bool, error)
}
