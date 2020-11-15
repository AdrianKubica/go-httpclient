package gohttp

import (
	"net/http"

	"github.com/AdrianKubica/go-httpclient/gomime"
)

func getHeaders(headers ...http.Header) http.Header {
	if len(headers) > 0 {
		return headers[0]
	}
	return http.Header{}
}

func (c *httpClient) getRequestHeader(header http.Header) http.Header {
	h := make(http.Header)

	// Add common headers to the request
	for k, v := range c.builder.header {
		if len(v) > 0 {
			h.Set(k, v[0])
		}
	}

	// Add custom headers to the request
	for k, v := range header {
		if len(v) > 0 {
			h.Set(k, v[0])
		}
	}

	// check if we actually have user agent defined in our headers if not set to passed value in builder
	if c.builder.userAgent != "" {
		if h.Get(gomime.HeaderUserAgent) != "" {
			return h
		}
		h.Set(gomime.HeaderUserAgent, c.builder.userAgent)
	}

	return h
}
