package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"github.com/project-nova/backend/pkg/logger"
)

const (
	// DefaultRetryCount represents how many time to retry a HTTP request by default.
	// We will start with no retry.
	DefaultRetryCount = 0

	// DefaultHTTPTimeout is the timeout for HTTP requests.
	DefaultHTTPTimeout = 60 * time.Second
)

// Make http easy to mock for unit testing
type Client interface {
	Request(method string, url string, params, result interface{}) (res *http.Response, err error)
	RequestAddHeaders(method string, url string,
		headers *map[string]string, params, result interface{}) (res *http.Response, err error)
}

// Client wraps a HTTP client.
type httpClient struct {
	BaseURL    string
	RetryCount int
	Timeout    time.Duration
	HTTPClient *http.Client
}

// ClientConfig contains configurations to construct a client.
type ClientConfig struct {
	BaseURL    string
	RetryCount int
	Timeout    time.Duration
	// Key        string
	// Passphrase string
	// Secret     string
}

// NewClient creates a new HTTP client.
func NewClient(config *ClientConfig) *httpClient {
	timeout := config.Timeout
	if timeout == 0 {
		timeout = DefaultHTTPTimeout
	}
	retryCount := config.RetryCount
	if retryCount == 0 {
		retryCount = DefaultRetryCount
	}
	client := httpClient{
		BaseURL:    config.BaseURL,
		RetryCount: retryCount,
		HTTPClient: &http.Client{
			Timeout: timeout,
		},
	}

	return &client
}

// Request issues a request to the server, and stores the response into result.
func (c *httpClient) Request(method string, url string, params, result interface{}) (res *http.Response, err error) {
	for i := 0; i < c.RetryCount+1; i++ {
		// TODO: use exponential backoff retry library instread of the below.
		retryDuration := time.Duration((math.Pow(2, float64(i))-1)/2*1000) * time.Millisecond
		time.Sleep(retryDuration)

		res, err = c.request(method, url, nil, params, result)
		if res != nil && res.StatusCode == 429 {
			logger.Warnf("HTTP 429 for %s %s: too many requests, rate limit hit", method, c.BaseURL+url)
			continue
		}
		logger.Debugf("request url %s method %s result: %s %v", c.BaseURL+url, method, result)
		break
	}

	return res, err
}

func (c *httpClient) RequestAddHeaders(method string, url string,
	headers *map[string]string, params, result interface{}) (res *http.Response, err error) {
	for i := 0; i < c.RetryCount+1; i++ {
		retryDuration := time.Duration((math.Pow(2, float64(i))-1)/2*1000) * time.Millisecond
		time.Sleep(retryDuration)

		res, err = c.request(method, url, headers, params, result)
		if res != nil && res.StatusCode == 429 {
			logger.Warnf("HTTP 429 for %s %s: too many requests, rate limit hit", method, c.BaseURL+url)
			continue
		}
		logger.Debugf("request url %s method %s result: %s %v", c.BaseURL+url, method, result)
		break
	}
	return res, err
}

func (c *httpClient) request(method string, url string,
	headers *map[string]string, params, result interface{}) (res *http.Response, err error) {
	var data []byte
	body := bytes.NewReader(make([]byte, 0))

	if params != nil {
		data, err = json.Marshal(params)
		if err != nil {
			return res, err
		}
		body = bytes.NewReader(data)
	}

	fullURL := fmt.Sprintf("%s%s", c.BaseURL, url)
	req, err := http.NewRequest(method, fullURL, body)
	if err != nil {
		return res, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	if headers != nil {
		for k, v := range *headers {
			req.Header.Add(k, v)
		}
	}
	logger.Debugf("HTTP request: %v", req)

	res, err = c.HTTPClient.Do(req)
	if err != nil {
		return res, errors.Wrapf(err, "HTTP request failure")
	}
	defer res.Body.Close()

	if res.StatusCode != 200 && res.StatusCode != 202 {
		defer res.Body.Close()
		e := Error{}
		decoder := json.NewDecoder(res.Body)
		if err := decoder.Decode(&e); err != nil {
			return res, errors.Wrapf(err, "failed to decode http response body error message")
		}

		return res, error(e)
	}

	if result != nil {
		decoder := json.NewDecoder(res.Body)
		if err = decoder.Decode(result); err != nil {
			return res, errors.Wrapf(err, "failed to decode http response into result: %v.\n%+v", res.Body, *res)
		}
	}

	return res, nil
}
