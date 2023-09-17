package repository

import "github.com/cocoide/tech-guide/pkg/domain/model"

//go:generate mockgen -source=comment.go -destination=../../../mock/repository/comment.go
type CommentRepo interface {
	CreateComment(comment *model.Comment) error
	GetComments(articleID int) ([]model.Comment, error)
}
