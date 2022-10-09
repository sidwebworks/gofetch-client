package gofetch

import (
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T) {

	// Initialization
	client := fetchClient{}

	customHeaderss := make(http.Header)
	customHeaderss.Set("Content-Type", "application/json")
	customHeaderss.Set("User-Agent", "gofetch-http-client")

	client.headers = customHeaderss

	// Execution
	requestHeaders := make(http.Header)

	requestHeaders.Set("X-Request-Id", "123-gofetch")

	fullHeaders := client.getRequestHeaders(requestHeaders)

	// Validation
	if len(fullHeaders) != 3 {
		t.Errorf("expected 3 headers but recieved: %v", len(fullHeaders))
	}

}
