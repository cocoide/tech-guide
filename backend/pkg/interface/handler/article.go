package handler

import (
	"context"
	"fmt"
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

func (h *Handler) GetRecommendArticles(c echo.Context) error {
	accountId := int(c.Get("account_id").(float64))
	strArticleIDs, err := h.cache.Get(fmt.Sprintf(key.PersonalizedArticleIDs, accountId))
	var articleIDs []int
	if len(strArticleIDs) < 1 || err != nil {
		articleIDs, err = h.personalize.GetRecommendArticleIDs(accountId)
		if err != nil {
			return c.JSON(400, err.Error())
		}
		strArticleIDs, err = utils.Serialize(articleIDs)
		if err != nil {
			return c.JSON(400, err.Error())
		}
		if err := h.cache.Set(fmt.Sprintf(key.PersonalizedArticleIDs, accountId), strArticleIDs, 24*time.Hour); err != nil {
			return c.JSON(400, err.Error())
		}
	} else {
		articleIDs, err = utils.Deserialize[[]int](strArticleIDs)
		if err != nil {
			return c.JSON(400, err.Error())
		}
	}
	articles, err := h.repo.GetArticlesByIDs(articleIDs)
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
	content, err := utils.GetMarkdownByURL(URL)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, content)
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
		if len(ogp.Thumbnail) > 700 {
			ogp.Thumbnail = ""
		}
		return c.JSON(200, ogp)
	}); err != nil {
		return c.JSON(401, err.Error())
	}
	return nil
}
func (h *Handler) GetSpeakerDeckID(c echo.Context) error {
	var result string
	URL := c.QueryParam("url")
	type Response struct {
		HTML string `json:"html"`
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
	articles, err := h.repo.GetLatestArticleByLimitWithSourceData(pageIndex, 20)
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
	if len(ogp.Thumbnail) < 500 {
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
		wholeTopicWeights, err := h.article.ExtractTopicsWithWeightFromArticleTitle(article.Title)
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
	articles, err := h.repo.GetArticlesBySourceID(sourceID, pageIndex, 10)
	return c.JSON(200, articles)
}

func (h *Handler) GetArticlesByTopicID(c echo.Context) error {
	topicID, err := strconv.Atoi(c.Param("topicId"))
	pageIndex := getPageIndexFromQuery(c, 1)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	articles, err := h.repo.GetArticlesByTopicID(topicID, pageIndex, 10)
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
