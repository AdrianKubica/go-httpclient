package examples

import (
	"net/http"
	"time"

	"github.com/AdrianKubica/go-httpclient/gohttp"
	"github.com/AdrianKubica/go-httpclient/gomime"
)

var (
	httpClient = getHttpClient()
)

func getHttpClient() gohttp.Client {
	headers := make(http.Header)
	headers.Set(gomime.HeaderContentType, gomime.ContentTypeJson)
	client := gohttp.NewBuilder().
		SetConnectionTimeout(2 * time.Second).
		SetResponseTimeout(3 * time.Second).
		SetHeader(headers).
		SetUserAgent("Webkit").
		Build()
	return client
}
