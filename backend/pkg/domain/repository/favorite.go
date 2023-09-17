package repository

//go:generate mockgen -source=favorite.go -destination=../../../mock/repository/favorite.go
type FavoriteRepo interface {
	DoFavoriteArticle(articleID, accountID int) error
	UnFavoriteArticle(articleID, accountID int) error
	CountFavoritesByArticleID(articleID int) (int64, error)
}
