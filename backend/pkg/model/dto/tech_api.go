package dto

import "time"

type QiitaArticleAPI struct {
	URL          string    `json:"url"`
	Title        string    `json:"title"`
	RenderedBody string    `json:"rendered_body"`
	CreatedAt    time.Time `json:"created_at"`
	LikesCount   int       `json:"likes_count"`
	StocksCount  int       `json:"stocks_count"`
	Tags         []struct {
		Name string `json:"name"`
	} `json:"tags"`
}
