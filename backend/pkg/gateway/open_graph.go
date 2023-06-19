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
	Image       string `json:"image_url"`
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
	var imageUrl string
	if ogp.Images != nil && len(ogp.Images) > 0 {
		imageUrl = ogp.Images[0].URL
	} else {
		imageUrl = ""
	}
	result := &OGP{
		Title:       ogp.Title,
		Image:       imageUrl,
		Description: ogp.Description,
	}
	return result, nil
}
