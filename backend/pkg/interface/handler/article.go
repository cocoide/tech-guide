package handler

import (
	"context"
	"fmt"
	"github.com/cocoide/tech-guide/pkg/interface/handler/ctxutils"
	"github.com/cocoide/tech-guide/pkg/usecase/parser"
	"log"
	"regexp"
	"strconv"
	"time"

	"github.com/cocoide/tech-guide/key"
	"github.com/cocoide/tech-guide/pkg/domain/model"
	"github.com/cocoide/tech-guide/pkg/domain/model/dto"
	"github.com/cocoide/tech-guide/pkg/utils"
	"github.com/labstack/echo"
)

const (
	paginateLimit = 6
)

func (h *Handler) GetRecommendArticles(c echo.Context) error {
	accountId := ctxutils.GetAccountID(c)
	strArticleIDs, exist, err := h.cache.Get(fmt.Sprintf(key.PersonalizedArticleIDs, accountId))
	var articleIDs []int
	if !exist || err != nil {
		articleIDs, err = h.personalize.GetRecommendArticleIDs(accountId)
		if err != nil {
			return c.JSON(400, err.Error())
		}
		if err != nil {
			return c.JSON(400, err.Error())
		}
		if err := h.cache.Set(fmt.Sprintf(key.PersonalizedArticleIDs, accountId), strArticleIDs, 24*time.Hour); err != nil {
			return c.JSON(400, err.Error())
		}
	} else {
		articleIDs, err = parser.Deserialize[[]int](string(strArticleIDs))
		if err != nil {
			return c.JSON(400, err.Error())
		}
	}
	articles, err := h.repo.GetArticlesByIDs(articleIDs, []string{"Source"})
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, articles)
}

func (h *Handler) GetRSS(c echo.Context) error {
	result, err := utils.FetchXML[dto.RSSFeed](c.QueryParam("url"), 5*time.Second)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, result)
}
func (h *Handler) GetDomainID(c echo.Context) error {
	URL := c.QueryParam("url")
	domain, err := utils.ExtractDomainNameFromURL(URL)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	id, err := h.repo.FindIDByDomain(domain)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, id)
}

func (h *Handler) GetOverview(c echo.Context) error {
	URL := c.QueryParam("url")
	headers, err := h.scraper.GetHeaders(URL)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	var strOverview string
	for _, header := range headers {
		strOverview += header.ToMarkdown() + "\n"
	}
	return c.JSON(200, strOverview)
}

func (h *Handler) GetOGP(c echo.Context) error {
	ctx := context.Background()
	var limiter = utils.NewRateLimiter(5, 10)
	if err := limiter.Limit(ctx, func() error {
		URL := c.QueryParam("url")
		ogp, err := h.ogp.GetOGPByURL(URL)
		if err != nil {
			return c.JSON(400, err.Error())
		}
		if len(ogp.Thumbnail) > 999 {
			ogp.Thumbnail = ""
		}
		return c.JSON(200, ogp)
	}); err != nil {
		return c.JSON(401, err.Error())
	}
	return nil
}

func getSpeakerDeckCacheKey(id string) string {
	return fmt.Sprintf("speakerdeck.%s", id)
}

func (h *Handler) GetSpeakerDeckID(c echo.Context) error {
	var result string
	URL := c.QueryParam("url")
	type Response struct {
		HTML string `json:"html"`
	}
	id, err := utils.ExtractIDFromURL(URL)
	if err != nil {
		return c.JSON(400, err)
	}
	value, exist, err := h.cache.Get(getSpeakerDeckCacheKey(id))
	if err != nil {
		return c.JSON(400, err)
	}
	if exist {
		result = value
		log.Printf("cache hit: %s", value)
		return c.JSON(200, result)
	}
	params := map[string]string{"url": URL}
	res, err := utils.FetchJSON[Response]("https://speakerdeck.com/oembed.json", params)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	pattern := `src="//speakerdeck.com/player/([^"]+)"`
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(res.HTML)
	if len(matches) > 1 {
		result = matches[1]
	} else {
		return c.JSON(400, "not found id")
	}
	h.cache.Set(getSpeakerDeckCacheKey(id), result, 30*24*time.Hour)
	return c.JSON(200, result)
}

func (h *Handler) GetArticles(c echo.Context) error {
	strPageIndex := c.QueryParam("page")
	var pageIndex int
	if strPageIndex == "" {
		pageIndex = 1
	} else {
		pageIndex, _ = strconv.Atoi(strPageIndex)
	}
	articles, err := h.repo.GetLatestArticleByLimitWithSourceData(pageIndex, paginateLimit)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, articles)
}

func (h *Handler) GetArticleDetail(c echo.Context) error {
	articleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err.Error())
	}
	articles, err := h.repo.GetArticleWithRelatedDataByID(articleID)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, articles)
}

func (h *Handler) GetRelatedArticles(c echo.Context) error {
	excludeArticleID := 0
	if c.QueryParam("exclude") != "" {
		excludeArticleID, _ = strconv.Atoi(c.QueryParam("exclude"))
	}
	originArticleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err.Error())
	}
	originTopicToArticleArray, err := h.repo.GetTopicToArticleArrayByArticleID(originArticleID)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	var result []model.Article
	if len(originTopicToArticleArray) > 0 {
		result, err = h.article.GetRelatedArticlesByOriginTopicToArticleArray(originTopicToArticleArray, excludeArticleID)
		if err != nil {
			return c.JSON(400, err.Error())
		}
	}
	return c.JSON(200, result)
}

func (h *Handler) CreateArticle(c echo.Context) error {
	type REQ struct {
		OriginalURL string `json:"original_url"`
	}
	req := new(REQ)
	if err := c.Bind(req); err != nil {
		return c.JSON(400, err.Error())
	}
	URL := req.OriginalURL
	ogp, err := h.ogp.GetOGPByURL(URL)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	thumbnailURL := ""
	if len(ogp.Thumbnail) < 999 {
		thumbnailURL = ogp.Thumbnail
	}
	article := &model.Article{
		Title:        ogp.Title,
		ThumbnailURL: thumbnailURL,
		OriginalURL:  URL,
	}
	if _, err := h.repo.CreateArticle(article); err != nil {
		return c.JSON(400, err.Error())
	}
	topicAssignErrCh := make(chan error)
	go func() {
		wholeTopicWeights, err := h.article.ExtractTopicsWithWeightFromArticleTitle(article.Title, article.Summary)
		if err != nil {
			topicAssignErrCh <- err
			return
		}
		topicWeights := dto.GetTopWeightedTopics(wholeTopicWeights, 3)
		var topicToArticles []model.TopicsToArticles
		for _, v := range topicWeights {
			topicToArticles = append(topicToArticles,
				model.TopicsToArticles{ArticleID: article.ID, TopicID: v.ID, Weight: v.Weight})
		}
		if err := h.repo.CreateTopicToArticle(topicToArticles); err != nil {
			topicAssignErrCh <- err
			return
		}
	}()
	go func() {
		for err := range topicAssignErrCh {
			log.Print(err)
		}
	}()
	return c.JSON(200, article)
}

func (h *Handler) GetArticlesBySourceID(c echo.Context) error {
	sourceID, err := strconv.Atoi(c.Param("sourceId"))
	pageIndex := getPageIndexFromQuery(c, 1)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	articles, err := h.repo.GetArticlesBySourceID(sourceID, pageIndex, paginateLimit)
	return c.JSON(200, articles)
}

func (h *Handler) GetArticlesByTopicID(c echo.Context) error {
	topicID, err := strconv.Atoi(c.Param("topicId"))
	pageIndex := getPageIndexFromQuery(c, 1)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	articles, err := h.repo.GetArticlesByTopicID(topicID, pageIndex, paginateLimit)
	return c.JSON(200, articles)
}

func getPageIndexFromQuery(c echo.Context, defaultPage int) int {
	strPageIndex := c.QueryParam("page")
	if strPageIndex == "" {
		return defaultPage
	}
	pageIndex, err := strconv.Atoi(strPageIndex)
	if err != nil {
		return defaultPage
	}
	return pageIndex
}

func (h *Handler) GetHeaders(c echo.Context) error {
	url := c.QueryParam("url")
	if len(url) < 1 {
		return c.JSON(400, "Required quey param: url")
	}
	headers, err := h.scraper.GetHeaders(url)
	if err != nil {
		return c.JSON(400, err)
	}

	var result string

	for _, header := range headers {
		result = result + header.ToMarkdown() + "\n"
	}
	return c.JSON(200, headers)
}

func (h *Handler) GetPageContent(c echo.Context) error {
	url := c.QueryParam("url")
	if len(url) < 1 {
		return c.JSON(400, "Required quey param: url")
	}
	level, _ := strconv.Atoi(c.QueryParam("level"))

	content, _ := h.scraper.GetHeaderContentWithHTMLElements(url, c.QueryParam("header"), level)

	return c.JSON(200, content)
}

func (h *Handler) FindByTitle(c echo.Context) error {
	title := c.QueryParam("title")
	articles, err := h.repo.FindArticlesByTitle(title)
	if err != nil {
		return c.JSON(400, err)
	}
	return c.JSON(200, articles)
}

func (h *Handler) Summarize(c echo.Context) error {
	url := c.QueryParam("url")
	if len(url) < 1 {
		return c.JSON(400, "Required quey param: url")
	}
	summary, err := h.scraping.SummarizeArticle(url)

	if err != nil {
		return c.JSON(400, err)
	}

	return c.JSON(200, summary)
}

func (h *Handler) ScrapeDetail(c echo.Context) error {
	url := c.QueryParam("url")
	content := c.QueryParam("content")
	level, _ := strconv.Atoi(c.QueryParam("level"))
	if len(url) < 1 {
		return c.JSON(400, "Required quey param: url")
	}
	summary, err := h.scraper.GetHeaderContent(url, content, level)

	if err != nil {
		return c.JSON(400, err)
	}

	return c.JSON(200, summary)
}
