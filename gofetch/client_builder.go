package gofetch

import (
	"net/http"
	"time"
)

type ClientBuilder interface {
	SetHeaders(headers http.Header) ClientBuilder
	SetConnectionTimeout(timeout time.Duration) ClientBuilder
	SetResponseTimeout(timeout time.Duration) ClientBuilder
	SetMaxIdleConnections(count int) ClientBuilder
	DisableTimeouts(disable bool) ClientBuilder
	Build() Client
}

type clientBuilder struct {
	maxIdleConnections int
	connectionTimeout  time.Duration
	responseTimeout    time.Duration
	disabledTimeouts   bool

	headers http.Header
}

func NewBuilder() ClientBuilder {
	builder := &clientBuilder{}

	return builder
}

func (c *clientBuilder) Build() Client {
	client := fetchClient{
		maxIdleConnections: c.maxIdleConnections,
		headers:            c.headers,
		connectionTimeout:  c.connectionTimeout,
		responseTimeout:    c.responseTimeout,
		disabledTimeouts:   c.disabledTimeouts,
	}

	return &client
}

func (c *clientBuilder) SetConnectionTimeout(timeout time.Duration) ClientBuilder {
	c.connectionTimeout = timeout

	return c
}

func (c *clientBuilder) SetResponseTimeout(timeout time.Duration) ClientBuilder {
	c.responseTimeout = timeout

	return c
}

func (c *clientBuilder) SetMaxIdleConnections(count int) ClientBuilder {
	c.maxIdleConnections = count

	return c
}

func (c *clientBuilder) SetHeaders(headers http.Header) ClientBuilder {
	c.headers = headers

	return c
}

func (c *clientBuilder) DisableTimeouts(disable bool) ClientBuilder {
	c.disabledTimeouts = disable

	return c
}
