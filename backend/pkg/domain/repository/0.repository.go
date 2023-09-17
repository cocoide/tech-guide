package repository

import (
	"context"
)

type Repository interface {
	AccountRepo
	TopicRepo
	ArticleRepo
	ActivityRepo
	CollectionRepo
	CommentRepo
	FavoriteRepo
	SourceRepo
	Transaction
}

type Transaction interface {
	DoInTx(ctx context.Context, fn func(ctx context.Context) error) error
}
