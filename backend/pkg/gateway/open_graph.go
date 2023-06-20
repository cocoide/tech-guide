package gateway

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/dyatlov/go-opengraph/opengraph"
)

type OGPGateway struct {
	client *http.Client
}

func NewOGPGateway() OGPGateway {
	return OGPGateway{client: &http.Client{Timeout: 10 * time.Second}}
}

type OGP struct {
	Title       string `json:"title"`
	Thumbnail   string `json:"thumbnail"`
	Sitename    string `json:"sitename"`
	Description string `json:"description"`
}

func (g *OGPGateway) GetOGPByURL(url string) (*OGP, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := g.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ogp := opengraph.NewOpenGraph()
	if err := ogp.ProcessHTML(strings.NewReader(string(body))); err != nil {
		return nil, err
	}
	var thumbnail string
	for _, image := range ogp.Images {
		if isImageValid(image.URL) && len(image.URL) < 1000 {
			thumbnail = image.URL
			break
		}
	}
	result := &OGP{
		Title:       ogp.Title,
		Thumbnail:   thumbnail,
		Sitename:    ogp.SiteName,
		Description: ogp.Description,
	}
	return result, nil
}

func isImageValid(imagePath string) bool {
	resp, err := http.Get(imagePath)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		return true
	}
	return false
}
