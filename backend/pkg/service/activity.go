package service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/cocoide/tech-guide/key"
	"github.com/cocoide/tech-guide/pkg/model"
	repo "github.com/cocoide/tech-guide/pkg/repository"
)

type ActivityService interface {
	AddContribution(accountId int, points int) error
	GetContributionsFromCache() ([]*model.Contribution, error)
}

type activityService struct {
	cr repo.CacheRepo
}

func NewActivityService(cr repo.CacheRepo) ActivityService {
	return &activityService{cr: cr}
}

func (s *activityService) AddContribution(accountId int, points int) error {
	strAccountId := strconv.Itoa(accountId)
	today := time.Now().Format("2006-01-02")
	key := fmt.Sprintf(key.Contributions, today)
	strContribute, err := s.cr.GetHashField(key, strAccountId)
	if err != nil {
		return err
	}
	if len(strContribute) > 0 {
		contribute, err := strconv.Atoi(strContribute)
		if err != nil {
			return err
		}
		field := map[string]interface{}{
			strAccountId: contribute + points,
		}
		if err := s.cr.SetHashField(key, field, 2*24*time.Hour); err != nil {
			return err
		}
	} else {
		field := map[string]interface{}{
			strAccountId: points,
		}
		if err := s.cr.SetHashField(key, field, 2*24*time.Hour); err != nil {
			return err
		}
	}
	return nil
}

func (s *activityService) GetContributionsFromCache() ([]*model.Contribution, error) {
	today := time.Now().Format("2006-01-02")
	key := fmt.Sprintf(key.Contributions, today)
	contributeMap, err := s.cr.GetAllHashFields(key)
	if err != nil {
		return nil, err
	}
	var contributions []*model.Contribution
	for field, value := range contributeMap {
		accountId, err := strconv.Atoi(field)
		if err != nil {
			continue
		}
		points, err := strconv.Atoi(value)
		if err != nil {
			continue
		}
		contributions = append(contributions, &model.Contribution{AccountID: accountId, Points: points, Date: time.Now()})
	}
	return contributions, nil
}
