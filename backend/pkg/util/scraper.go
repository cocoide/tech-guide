package util

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GetMarkdownByURL(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return "", fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}

	var content strings.Builder

	doc.Find("h2, h3, h4").Each(func(i int, s *goquery.Selection) {
		headerLevel := 0
		switch goquery.NodeName(s) {
		case "h2":
			headerLevel = 2
		case "h3":
			headerLevel = 3
		case "h4":
			headerLevel = 4
		}
		for i := 0; i < headerLevel; i++ {
			content.WriteString("#")
		}
		if headerLevel >= 2 && headerLevel <= 4 {
			content.WriteString(" ")
			content.WriteString(s.Text())
			content.WriteString("\n\n")
		}
	})

	return content.String(), nil
}
