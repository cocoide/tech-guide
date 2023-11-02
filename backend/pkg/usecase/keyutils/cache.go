package keyutils

import "fmt"

func CacheFeedArticleIDs(accountId int) string {
	return fmt.Sprintf("feed_article_ids.%d", accountId)
}
