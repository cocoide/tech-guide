package scheduler

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type HttpMethod int

const (
	GET HttpMethod = iota + 1
	POST
	DELTE
	PUT
)

type ContentType string

const (
	JSON      ContentType = "application/json"
	PLAIN     ContentType = "text/plain"
	FORM_DATA ContentType = "multipart/form-data"
)

type HttpClient struct {
	Client   *http.Client
	Endpoint string
	Headers  map[string]string
	Params   map[string]interface{}
	Body     io.Reader
}

func NewHttpClient() *HttpClient {
	return &HttpClient{
		Client:   &http.Client{Timeout: 10 * time.Second},
		Endpoint: "",
		Headers:  make(map[string]string),
		Params:   make(map[string]interface{}),
		Body:     nil,
	}
}

// defaultの10秒を上書きできる
func (h *HttpClient) WithTimeout(duration time.Duration) *HttpClient {
	h.Client.Timeout = duration
	return h
}

func (h *HttpClient) IsTimeoutError(err error) bool {
	var urlErr *url.Error
	if errors.As(err, &urlErr) {
		return urlErr.Timeout()
	}
	return false
}

func (h *HttpClient) WithBaseURL(baseURL string) *HttpClient {
	h.Endpoint = baseURL
	return h
}

func (h *HttpClient) WithBody(values []byte, content ContentType) *HttpClient {
	h.Body = bytes.NewReader(values)
	h.Headers["Content-Type"] = string(content)
	return h
}

func (h *HttpClient) WithBearerToken(token string) *HttpClient {
	h.Headers["Authorization"] = fmt.Sprintf("Bearer %s", token)
	return h
}

func (h *HttpClient) WithPath(path string) *HttpClient {
	h.Endpoint = h.Endpoint + "/" + path
	return h
}

func (h *HttpClient) WithHeader(key, value string) *HttpClient {
	h.Headers[key] = value
	return h
}

func (h *HttpClient) WithParam(key string, value interface{}) *HttpClient {
	h.Params[key] = value
	return h
}

func (h *HttpClient) WithRawParams(rawQueryParams map[string]interface{}) *HttpClient {
	var params []string
	for key, value := range rawQueryParams {
		params = append(params, fmt.Sprintf("%s=%v", key, value))
	}
	joinedParams := strings.Join(params, "&")
	h.Endpoint = h.Endpoint + "?" + joinedParams
	return h
}

func (h *HttpClient) ExecuteRequest(method HttpMethod) ([]byte, error) {
	var methodName string
	switch method {
	case GET:
		methodName = "GET"
	case POST:
		methodName = "POST"
	case DELTE:
		methodName = "DELETE"
	case PUT:
		methodName = "PUT"
	}
	client := h.Client

	req, err := http.NewRequest(methodName, h.Endpoint, h.Body)
	if err != nil {
		return nil, err
	}

	for k, v := range h.Headers {
		req.Header.Add(k, v)
	}

	query := req.URL.Query()
	for key, value := range h.Params {
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

	statusCode := resp.StatusCode

	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("Request failed with status code: %d", statusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
