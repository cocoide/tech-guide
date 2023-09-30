package main

import (
	"fmt"
	"github.com/cocoide/tech-guide/pkg/domain/model"
	"github.com/cocoide/tech-guide/pkg/infrastructure/database"
	"github.com/cocoide/tech-guide/pkg/infrastructure/integration"
	"log"
	"sync"
)

func main() {
	db := database.NewGormDatabase()
	var articles []model.Article
	idToThumbnailMap := make(map[int]string, len(articles))
	ogp := integration.NewOGPService()

	if err := db.Find(&articles).Where("thumbnail_url = ?", "").Error; err != nil {
		log.Print(err)
	}
	var wg sync.WaitGroup
	var mu sync.Mutex
	batchSize := 10
	for i := 0; i < len(articles); i += batchSize {
		end := i + batchSize
		if end > len(articles) {
			end = len(articles)
		}

		currentBatch := articles[i:end]
		for _, v := range currentBatch {
			wg.Add(1)
			go func(v model.Article) {
				defer wg.Done()
				ogpResponse, err := ogp.GetOGPByURL(v.OriginalURL)
				fmt.Printf("%sのOGP取得", v.OriginalURL)
				if err != nil {
					log.Print(err)
				}
				mu.Lock()
				idToThumbnailMap[v.ID] = ogpResponse.Thumbnail
				mu.Unlock()
			}(v)
		}
		wg.Wait()
	}
	//fmt.Printf("map: %v", idToThumbnailMap)
	ids := []int{}
	cases := ""
	for id, thumbnail := range idToThumbnailMap {
		cases += fmt.Sprintf("WHEN %d THEN '%s' ", id, thumbnail)
		ids = append(ids, id)
	}
	sql := fmt.Sprintf("UPDATE articles SET thumbnail_url = CASE id %s END", cases)
	err := db.Exec(sql).Error
	if err != nil {
		log.Print(err)
	}
}
