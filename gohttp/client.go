package gohttp

import (
	"net/http"
	"sync"

	"github.com/AdrianKubica/go-httpclient/core"
)

type httpClient struct {
	// composition, embedded struct
	builder *clientBuilder
	// defaults to nil if not initialized with value and then cause panic
	client *http.Client
	// sync.Once defeats from creating many instances of client by goroutines
	clientOnce sync.Once
}

type Client interface {
	Get(url string, headers ...http.Header) (*core.Response, error)
	Post(url string, body interface{}, headers ...http.Header) (*core.Response, error)
	Put(url string, body interface{}, headers ...http.Header) (*core.Response, error)
	Patch(url string, body interface{}, headers ...http.Header) (*core.Response, error)
	Delete(url string, headers ...http.Header) (*core.Response, error)
	Options(url string, headers ...http.Header) (*core.Response, error)
}

func (c *httpClient) Get(url string, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodGet, url, getHeaders(headers...), nil)
}

func (c *httpClient) Post(url string, body interface{}, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodPost, url, getHeaders(headers...), body)
}

func (c *httpClient) Put(url string, body interface{}, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodPut, url, getHeaders(headers...), body)
}

func (c *httpClient) Patch(url string, body interface{}, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodPatch, url, getHeaders(headers...), body)
}

func (c *httpClient) Delete(url string, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodDelete, url, getHeaders(headers...), nil)
}

func (c *httpClient) Options(url string, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodOptions, url, getHeaders(headers...), nil)
}
