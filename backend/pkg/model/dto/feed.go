package dto

import (
	"encoding/xml"

	"github.com/cocoide/tech-guide/pkg/model"
)

type RSSFeed struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	Image       struct {
		URL  string `xml:"url"`
		Link string `xml:"link"`
	} `xml:"image"`
	LastBuildDate string `xml:"lastBuildDate"`
	Language      string `xml:"language"`
	Items         []struct {
		Title       string `xml:"title"`
		Description string `xml:"description"`
		Link        string `xml:"link"`
		Enclosure   struct {
			URL    string `xml:"url,attr"`
			Length string `xml:"length,attr"`
			Type   string `xml:"type,attr"`
		} `xml:"enclosure"`
		Creator string `xml:"creator"`
		Content string `xml:"encoded"`
	} `xml:"item"`
}

func (z *RSSFeed) GetArticles() ([]model.Article, error) {
	articles := make([]model.Article, len(z.Channel.Items))
	for i, item := range z.Channel.Items {
		article := model.Article{
			Title:        item.Title,
			ThumbnailURL: item.Enclosure.URL,
			OriginalURL:  item.Link,
		}
		articles[i] = article
	}
	return articles, nil
}
