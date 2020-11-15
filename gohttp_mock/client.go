package gohttp_mock

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type httpClientMock struct{}

func (c *httpClientMock) Do(req *http.Request) (*http.Response, error) {
	// use of req.GetBody which is a req.Body copy otherwise you will close main body from furthermore read
	reqBody, err := req.GetBody()
	if err != nil {
		return nil, err
	}
	defer reqBody.Close()

	body, err := ioutil.ReadAll(reqBody)
	if err != nil {
		return nil, err
	}

	var res http.Response

	mock := MockupServer.mocks[MockupServer.getMockKey(req.Method, req.URL.String(), string(body))]
	if mock != nil {
		if mock.Error != nil {
			return nil, mock.Error
		}
		res.StatusCode = mock.ResponseStatusCode
		res.Body = ioutil.NopCloser(strings.NewReader(mock.ResponseBody))
		res.ContentLength = int64(len(mock.ResponseBody))
		res.Request = req
		return &res, nil
	}
	return nil, errors.New(fmt.Sprintf("no mock matching %s from '%s' with given body", req.Method, req.URL.String()))
}
