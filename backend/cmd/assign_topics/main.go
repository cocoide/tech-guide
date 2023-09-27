package assign_topics

import (
	"github.com/cocoide/tech-guide/pkg/infrastructure/database"
	"github.com/cocoide/tech-guide/pkg/infrastructure/rdb"
)

func main() {
	db := database.NewGormDatabase()
	repo := rdb.NewRepository(db)

	repo.GetArtucke
	repo.CreateTopicToArticle()
	repo.DoInTx()
}
