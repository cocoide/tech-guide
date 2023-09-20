package dto

import (
	"github.com/cocoide/tech-guide/pkg/domain/model"
	"time"
)

type GithubRepo struct {
	Name            string
	Owner           string
	StarCount       int
	PrimaryLanguage string
	Description     string
	HomepageURL     string
	CreatedAt       time.Time
	Topics          []model.Topic
}
