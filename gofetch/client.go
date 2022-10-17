package gofetch

import (
	"net"
	"net/http"
	"sync"
	"time"
)

type Client interface {
	Get(url string, headers http.Header) (*Response, error)
	Post(url string, headers http.Header, body interface{}) (*Response, error)
	Patch(url string, headers http.Header, body interface{}) (*Response, error)
	Put(url string, headers http.Header, body interface{}) (*Response, error)
	Delete(url string, headers http.Header) (*Response, error)
	Head(url string, headers http.Header) (*Response, error)
}

type fetchClient struct {
	client   *http.Client
	builder  *clientBuilder
	syncOnce sync.Once
}

const (
	defaultMaxIdleConnections = 5
	defaultResponseTimeout    = 5 * time.Second
	defaultConnectionTimeout  = 1 * time.Second
)

func (c *fetchClient) getHttpClient() *http.Client {
	c.syncOnce.Do(func() {
		hc := &http.Client{
			Timeout: c.getConnectionTimeout() + c.getResponseTimeout(),
			Transport: &http.Transport{
				MaxIdleConnsPerHost:   c.getMaxIdleConnections(),
				ResponseHeaderTimeout: c.getResponseTimeout(),
				DialContext: (&net.Dialer{
					Timeout: c.getConnectionTimeout(),
				}).DialContext,
			},
		}

		c.client = hc
	})

	return c.client
}

func (c *fetchClient) Get(url string, headers http.Header) (*Response, error) {
	return c.request(http.MethodGet, url, headers, nil)
}

func (c *fetchClient) Post(url string, headers http.Header, body interface{}) (*Response, error) {
	return c.request(http.MethodPost, url, headers, body)
}

func (c *fetchClient) Patch(url string, headers http.Header, body interface{}) (*Response, error) {
	return c.request(http.MethodPatch, url, headers, body)
}

func (c *fetchClient) Put(url string, headers http.Header, body interface{}) (*Response, error) {
	return c.request(http.MethodPut, url, headers, body)
}

func (c *fetchClient) Delete(url string, headers http.Header) (*Response, error) {
	return c.request(http.MethodDelete, url, headers, nil)
}

func (c *fetchClient) Head(url string, headers http.Header) (*Response, error) {
	return c.request(http.MethodHead, url, headers, nil)
}

func (c *fetchClient) getMaxIdleConnections() int {
	if c.builder.maxIdleConnections > 0 {
		return c.builder.maxIdleConnections
	}

	return defaultMaxIdleConnections
}

func (c *fetchClient) getConnectionTimeout() time.Duration {

	if c.builder.disabledTimeouts {
		return 0
	}

	if c.builder.connectionTimeout > 0 {
		return c.builder.connectionTimeout
	}

	return defaultConnectionTimeout
}

func (c *fetchClient) getResponseTimeout() time.Duration {
	if c.builder.responseTimeout > 0 {
		return c.builder.responseTimeout
	}

	if c.builder.disabledTimeouts {
		return 0
	}

	return defaultResponseTimeout
}
