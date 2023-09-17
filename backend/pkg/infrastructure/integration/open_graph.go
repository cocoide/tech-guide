package integration

import (
	"io"
	"net/http"
	"strings"

	"github.com/cocoide/tech-guide/pkg/domain/model/dto"
	"github.com/dyatlov/go-opengraph/opengraph"
)

func (s *ogpService) GetOGPByURL(url string) (*dto.OGPResponse, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ogp := opengraph.NewOpenGraph()
	if err := ogp.ProcessHTML(strings.NewReader(string(body))); err != nil {
		return nil, err
	}
	var thumbnail string
	for _, image := range ogp.Images {
		if isImageValid(image.URL) {
			thumbnail = image.URL
			break
		}
	}
	result := &dto.OGPResponse{
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
