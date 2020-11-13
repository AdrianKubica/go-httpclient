package gohttp

import (
	"net/http"
	"sync"
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
	Get(url string, header http.Header) (*Response, error)
	Post(url string, header http.Header, body interface{}) (*Response, error)
	Put(url string, header http.Header, body interface{}) (*Response, error)
	Patch(url string, header http.Header, body interface{}) (*Response, error)
	Delete(url string, header http.Header) (*Response, error)
}

func (c *httpClient) Get(url string, header http.Header) (*Response, error) {
	return c.do(http.MethodGet, url, header, nil)
}

func (c *httpClient) Post(url string, header http.Header, body interface{}) (*Response, error) {
	return c.do(http.MethodPost, url, header, body)
}

func (c *httpClient) Put(url string, header http.Header, body interface{}) (*Response, error) {
	return c.do(http.MethodPut, url, header, body)
}

func (c *httpClient) Patch(url string, header http.Header, body interface{}) (*Response, error) {
	return c.do(http.MethodPatch, url, header, body)
}

func (c *httpClient) Delete(url string, header http.Header) (*Response, error) {
	return c.do(http.MethodDelete, url, header, nil)
}
