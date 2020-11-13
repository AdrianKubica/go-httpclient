package gohttp

import (
	"fmt"
	"net/http"
)

type Mock struct {
	Method      string
	Url         string
	RequestBody string

	ResponseBody       string
	Error              error
	ResponseStatusCode int
}

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
