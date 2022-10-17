package gofetch

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type Response struct {
	status     string
	statusCode int
	body       []byte
	headers    http.Header
}

func (r *Response) Status() string {
	return r.status
}

func (r *Response) StatusCode() int {
	return r.statusCode
}

func (r *Response) Headers() http.Header {
	return r.headers
}

func (r *Response) Json(target interface{}) error {
	return json.Unmarshal(r.Bytes(), target)
}

func (r *Response) Text() string {
	return string(r.body)
}

func (r *Response) Bytes() []byte {
	return r.body
}

func (c *fetchClient) request(method string, url string, headers http.Header, body interface{}) (*Response, error) {

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

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	response := Response{
		status:     res.Status,
		statusCode: res.StatusCode,
		body:       bytes,
		headers:    res.Header,
	}

	return &response, nil

}

func (c *fetchClient) getRequestHeaders(customHeaders http.Header) http.Header {
	result := make(http.Header)

	for header, value := range c.builder.headers {
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
