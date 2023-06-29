package promptlayer

import (
	"net/http"
)

// MockHTTPClient is a mock implementation of the HTTPClient interface.
type mockHTTPClient struct {
	doFunc func(req *http.Request) (*http.Response, error)
}

// Do performs the mock HTTP request.
func (m *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return m.doFunc(req)
}
