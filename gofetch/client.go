package gofetch

import (
	"net"
	"net/http"
	"time"
)

type Client interface {
	Get(url string, headers http.Header) (*http.Response, error)
	Post(url string, headers http.Header, body interface{}) (*http.Response, error)
	Patch(url string, headers http.Header, body interface{}) (*http.Response, error)
	Put(url string, headers http.Header, body interface{}) (*http.Response, error)
	Delete(url string, headers http.Header) (*http.Response, error)
	Head(url string, headers http.Header) (*http.Response, error)
}

type fetchClient struct {
	client *http.Client

	// Client options
	maxIdleConnections int
	connectionTimeout  time.Duration
	responseTimeout    time.Duration
	disabledTimeouts   bool

	headers http.Header
}

const (
	defaultMaxIdleConnections = 5
	defaultResponseTimeout    = 5 * time.Second
	defaultConnectionTimeout  = 1 * time.Second
)

func New() Client {

	c := &fetchClient{}

	return c
}

func (c *fetchClient) getHttpClient() *http.Client {
	if c.client != nil {
		return c.client
	}

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

	return hc
}

func (c *fetchClient) Get(url string, headers http.Header) (*http.Response, error) {
	return c.request(http.MethodGet, url, headers, nil)
}

func (c *fetchClient) Post(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.request(http.MethodPost, url, headers, body)
}

func (c *fetchClient) Patch(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.request(http.MethodPatch, url, headers, body)
}

func (c *fetchClient) Put(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.request(http.MethodPut, url, headers, body)
}

func (c *fetchClient) Delete(url string, headers http.Header) (*http.Response, error) {
	return c.request(http.MethodDelete, url, headers, nil)
}

func (c *fetchClient) Head(url string, headers http.Header) (*http.Response, error) {
	return c.request(http.MethodHead, url, headers, nil)
}

func (c *fetchClient) getMaxIdleConnections() int {
	if c.maxIdleConnections > 0 {
		return c.maxIdleConnections
	}

	return defaultMaxIdleConnections
}

func (c *fetchClient) getConnectionTimeout() time.Duration {
	if c.connectionTimeout > 0 {
		return c.connectionTimeout
	}

	if c.disabledTimeouts {
		return 0
	}

	return defaultConnectionTimeout
}

func (c *fetchClient) getResponseTimeout() time.Duration {
	if c.responseTimeout > 0 {
		return c.responseTimeout
	}

	if c.disabledTimeouts {
		return 0
	}

	return defaultResponseTimeout
}
