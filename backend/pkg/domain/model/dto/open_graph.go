package dto

type OGPResponse struct {
	URL         string `json:"url"`
	Title       string `json:"title"`
	Thumbnail   string `json:"thumbnail"`
	Sitename    string `json:"sitename"`
	Description string `json:"description"`
}
