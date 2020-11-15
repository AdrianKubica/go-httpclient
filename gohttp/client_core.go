package gohttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/AdrianKubica/go-httpclient/core"
	"github.com/AdrianKubica/go-httpclient/gohttp_mock"
	"github.com/AdrianKubica/go-httpclient/gomime"
)

const (
	defaultMaxIdleConnections = 5
	defaultResponseTimeout    = 5 * time.Second
	defaultConnectionTimeout  = 1 * time.Second
)

func (c *httpClient) do(method string, url string, header http.Header, body interface{}) (*core.Response, error) {
	h := c.getRequestHeader(header)

	b, err := c.getRequestBody(h.Get("Content-Type"), body)
	if err != nil {
		return nil, err
	}

	if mock := gohttp_mock.GetMock(method, url, string(b)); mock != nil {
		return mock.GetResponse()
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(b))
	if err != nil {
		return nil, errors.New("unable to create a new request")
	}

	req.Header = h

	client := c.getHttpClient()

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	resFinal := core.Response{
		Status:     res.Status,
		StatusCode: res.StatusCode,
		Headers:    res.Header,
		Body:       resBody,
	}
	return &resFinal, nil
}

func (c *httpClient) getHttpClient() *http.Client {
	c.clientOnce.Do(func() {
		if c.builder.client != nil {
			c.client = c.builder.client
			return
		}
		c.client = &http.Client{
			Timeout: c.getConnectionTimeout() + c.getResponseTimeout(),
			// Transport property of http.Client is RoundTripper which is an interface so we need a pointer
			Transport: &http.Transport{
				MaxIdleConnsPerHost: c.getMaxIdleConnections(),
				// max amount of time to wait for getting response header from server when request is send
				// (time is analyzed while header is retrieved from server, body read is excluded, request timeout)
				ResponseHeaderTimeout: c.getResponseTimeout(),
				DialContext: (&net.Dialer{
					// max amount of time to wait for a given connection, connection timeout
					Timeout: c.getConnectionTimeout(),
				}).DialContext,
			},
		}
	})
	return c.client
}

func (c *httpClient) getMaxIdleConnections() int {
	if c.builder.maxIdleConnections > 0 {
		return c.builder.maxIdleConnections
	}
	return defaultMaxIdleConnections
}

func (c *httpClient) getResponseTimeout() time.Duration {
	if c.builder.disableTimeouts {
		return 0
	}
	if c.builder.responseTimeout > 0 {
		return c.builder.responseTimeout
	}
	return defaultResponseTimeout
}

func (c *httpClient) getConnectionTimeout() time.Duration {
	if c.builder.connectionTimeout > 0 {
		return c.builder.connectionTimeout
	}
	if c.builder.disableTimeouts {
		return 0
	}
	return defaultConnectionTimeout
}

func (c *httpClient) getRequestBody(contentType string, body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}

	switch strings.ToLower(contentType) {
	case gomime.ContentTypeJson:
		return json.Marshal(body)
	case gomime.ContentTypeXML:
		return xml.Marshal(body)
	default:
		return json.Marshal(body)
	}
}
