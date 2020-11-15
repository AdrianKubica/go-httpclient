package gohttp_mock

import (
	"fmt"
	"net/http"

	"github.com/AdrianKubica/go-httpclient/core"
)

// Mock provides a way to configure HTTP mocks based on combination
// between request method, URL and request body
type Mock struct {
	Method      string
	Url         string
	RequestBody string

	ResponseBody       string
	Error              error
	ResponseStatusCode int
}

// GetResponse returns a Response struct based on the mock configuration
func (m *Mock) GetResponse() (*core.Response, error) {
	if m.Error != nil {
		return nil, m.Error
	}

	res := core.Response{
		Status:     fmt.Sprintf("%d %s", m.ResponseStatusCode, http.StatusText(m.ResponseStatusCode)),
		StatusCode: m.ResponseStatusCode,
		Body:       []byte(m.ResponseBody),
	}
	return &res, nil
}
