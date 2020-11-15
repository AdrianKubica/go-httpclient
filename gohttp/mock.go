package gohttp

import (
	"fmt"
	"net/http"
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
func (m *Mock) GetResponse() (*Response, error) {
	if m.Error != nil {
		return nil, m.Error
	}

	res := Response{
		status:     fmt.Sprintf("%d %s", m.ResponseStatusCode, http.StatusText(m.ResponseStatusCode)),
		statusCode: m.ResponseStatusCode,
		body:       []byte(m.ResponseBody),
	}
	return &res, nil
}
