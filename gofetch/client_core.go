package gofetch

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"net/http"
	"strings"
)

func (c *fetchClient) request(method string, url string, headers http.Header, body interface{}) (*http.Response, error) {

	requestHeaders := c.getRequestHeaders(headers)

	requestBody, err := c.getRequestBody(headers.Get("contentType"), body)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))

	if err != nil {
		return nil, errors.New("unable to create a new request")
	}

	req.Header = requestHeaders

	client := c.getHttpClient()

	return client.Do(req)
}

func (c *fetchClient) getRequestHeaders(customHeaders http.Header) http.Header {
	result := make(http.Header)

	for header, value := range c.headers {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	for header, value := range customHeaders {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	return result
}

func (c *fetchClient) getRequestBody(contentType string, body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}

	switch strings.ToLower(contentType) {
	case "application/json":
		return json.Marshal(body)
	case "application/xml":
		return xml.Marshal(body)
	default:
		return xml.Marshal(body)
	}
}
