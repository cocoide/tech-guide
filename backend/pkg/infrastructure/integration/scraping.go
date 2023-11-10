package integration

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/cocoide/tech-guide/pkg/domain/model/dto"
	"github.com/cocoide/tech-guide/pkg/domain/service"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type scrapingService struct {
}

func NewScrapingService() service.ScrapingService {
	return &scrapingService{}
}

func (s *scrapingService) GetHeaders(url string) ([]*dto.Header, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var headers []*dto.Header
	var currentH1 *dto.Header
	var currentH2 *dto.Header
	var currentH3 *dto.Header

	doc.Find("h1, h2, h3, h4").Each(func(i int, query *goquery.Selection) {
		level := 0
		switch goquery.NodeName(query) {
		case "h1":
			level = 1
		case "h2":
			level = 2
		case "h3":
			level = 3
		case "h4":
			level = 4
		}
		trimStrings := []string{"\n", " ", "\t"}
		content := s.removeLeadingStrings(query.Text(), trimStrings)
		newHeader := &dto.Header{
			Level:   level,
			Content: content,
		}

		switch level {
		case 1:
			headers = append(headers, newHeader)
			currentH1 = newHeader
			currentH2 = nil
			currentH3 = nil
		case 2:
			if currentH1 != nil {
				currentH1.AddSubHeader(newHeader)
			} else {
				headers = append(headers, newHeader) // treat as standalone if no h1 parent
			}
			currentH2 = newHeader
			currentH3 = nil // reset current h3
		case 3:
			if currentH2 != nil {
				currentH2.AddSubHeader(newHeader)
			} else if currentH1 != nil {
				currentH1.AddSubHeader(newHeader)
			} else {
				headers = append(headers, newHeader) // treat as standalone if no h2 parent
			}
			currentH3 = newHeader
		case 4:
			if currentH3 != nil {
				currentH3.AddSubHeader(newHeader)
			} else if currentH2 != nil {
				currentH2.AddSubHeader(newHeader)
			} else if currentH1 != nil {
				currentH1.AddSubHeader(newHeader)
			} else {
				headers = append(headers, newHeader) // treat as standalone if no h2/h3 parent
			}
		}
	})

	return headers, nil
}

func (s *scrapingService) GetAllContentByURL(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", fmt.Errorf("failed to fetch the URL with status code: %d", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}

	// HTMLドキュメント全体のテキストを取得する
	content := doc.Text()
	return content, nil
}

func (s *scrapingService) removeLeadingStrings(text string, srings []string) string {
	changed := true
	for changed {
		changed = false
		for _, sub := range srings {
			if strings.HasPrefix(text, sub) {
				text = strings.TrimPrefix(text, sub)
				changed = true
			}
		}
	}
	regexPattern := `^\d+\.\s*`
	re := regexp.MustCompile(regexPattern)
	text = re.ReplaceAllString(text, "")
	return text
}

func (s *scrapingService) GetHeaderContent(url string, headerContent string, headerLevel int) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return "", fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	var result string

	headerSelector := fmt.Sprintf("h%d", headerLevel)

	doc.Find(headerSelector).EachWithBreak(func(i int, h *goquery.Selection) bool {

		if strings.Contains(h.Text(), headerContent) {
			for node := h.Next(); node.Size() > 0; node = node.Next() {
				nodeName := goquery.NodeName(node)
				if level, err := strconv.Atoi(strings.TrimPrefix(nodeName, "h")); err == nil && level <= headerLevel {
					// 指定したヘッダーレベルまたはそれ以上のヘッダ要素に達したら停止
					break
				}
				result += node.Text()
			}
			return false
		}
		return true
	})

	return strings.TrimSpace(result), nil
}

func (s *scrapingService) GetHeaderContentWithHTMLElements(url string, headerContent string, headerLevel int) (string, error) {

	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return "", fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	var result string

	headerSelector := fmt.Sprintf("h%d", headerLevel)

	doc.Find(headerSelector).EachWithBreak(func(i int, h *goquery.Selection) bool {

		if strings.Contains(h.Text(), headerContent) {
			for node := h.Next(); node.Size() > 0; node = node.Next() {
				nodeName := goquery.NodeName(node)
				// 指定したヘッダーレベルまたはそれ以上のヘッダ要素に達したら停止
				if level, err := strconv.Atoi(strings.TrimPrefix(nodeName, "h")); err == nil && level <= headerLevel {
					break
				}
				node.Find("a").Each(func(j int, aNode *goquery.Selection) {
					href, exists := aNode.Attr("href")
					if !exists {
						return // hrefが存在しない場合、次のイテレーションに進む
					}
					isImageLink := false
					// ネストされたimgタグを取得
					aNode.Find("img").Each(func(k int, imgNode *goquery.Selection) {
						if src, srcExists := imgNode.Attr("src"); srcExists {
							alt := imgNode.AttrOr("alt", "")
							result += fmt.Sprintf("[![%s](%s)](%s)", alt, src, href)
							isImageLink = true
						}
					})
					currentElementName := goquery.NodeName(node)
					if !isImageLink && currentElementName == "h1" || currentElementName == "h2" || currentElementName == "h3" ||
						currentElementName == "h4" || currentElementName == "h5" || currentElementName == "h6" {
						result += aNode.Text()
					} else if !isImageLink {
						result += fmt.Sprintf("[%s](%s)", aNode.Text(), href)
					}
				})

				if goquery.NodeName(node) != "img" && goquery.NodeName(node) != "a" {
					result += node.Text()
				}
			}
			return false
		}
		return true
	})

	return strings.TrimSpace(result), nil
}
