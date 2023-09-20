package integration

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

type HttpClient struct {
	Client   *http.Client
	Endpoint string
	Headers  map[string]string
	Params   map[string]interface{}
}

func NewHttpClient() *HttpClient {
	return &HttpClient{
		Client:   &http.Client{Timeout: 10 * time.Second},
		Endpoint: "",
		Headers:  make(map[string]string),
		Params:   make(map[string]interface{}),
	}
}

func (h *HttpClient) WithHeader(key, value string) {
	h.Headers[key] = value
}

func (h *HttpClient) WithParam(key string, value interface{}) {
	h.Params[key] = value
}

func (h *HttpClient) WithBaseURL(baseURL string) {
	h.Endpoint = baseURL
}

func (h *HttpClient) WithPath(path string) {
	h.Endpoint = h.Endpoint + path
}

func (h *HttpClient) GetAPI() ([]byte, error) {
	client := h.Client

	req, err := http.NewRequest("GET", h.Endpoint, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range h.Headers {
		req.Header.Add(k, v)
	}

	query := req.URL.Query()
	for key, value := range h.Params {
		log.Print(value)
		switch v := value.(type) {
		case string:
			query.Add(key, v)
		case int:
			query.Add(key, strconv.Itoa(v))
		case bool:
			query.Add(key, strconv.FormatBool(v))
		default:
			return nil, fmt.Errorf("Failed to parse param value: %v", value)
		}
	}
	req.URL.RawQuery = query.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
