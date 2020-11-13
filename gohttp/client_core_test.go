package gohttp

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"testing"
)

func TestGetRequestHeader(t *testing.T) {
	t.Run("Check if headers are properly propagated", func(t *testing.T) {
		// TEST CASES
		headers := []struct {
			common   http.Header
			custom   http.Header
			expected http.Header
		}{
			{
				http.Header{"Content-Type": {"application/json"}, "User-Agent": {"cool-http-client"}},
				http.Header{"X-Request-Id": {"ABC-123"}},
				http.Header{"Content-Type": {"application/json"}, "User-Agent": {"cool-http-client"}, "X-Request-Id": {"ABC-123"}},
			},
			{
				http.Header{"Content-Type": {"application/json"}},
				http.Header{"X-Request-Id": {"ABC-123"}},
				http.Header{"Content-Type": {"application/json"}, "X-Request-Id": {"ABC-123"}},
			},
			{
				nil,
				nil,
				nil,
			},
		}

		for _, header := range headers {
			// INITIALIZATION
			client := httpClient{builder: &clientBuilder{header: header.common}}
			// EXECUTION
			actual := client.getRequestHeader(header.custom)
			// VALIDATION
			for k, exp := range header.expected {
				if act, ok := actual[k]; !ok {
					t.Errorf("There is no a header for key: %s with expected value %s got: %s", k, exp, act)
				}
				if act, ok := actual[k]; ok && act[0] != exp[0] {
					t.Errorf("Expected header for key: %s with expected value %s got: %s", k, exp, act)
				}
			}

			expLen := len(header.expected)
			actLen := len(actual)
			if expLen != actLen {
				t.Errorf("Header lengt inconsistency, expected: %d, got: %d", expLen, actLen)
			}
		}
	})
}

func TestGetRequestBody(t *testing.T) {
	// INITIALIZATION

	type Person struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Exp       int    `json:"experience"`
	}

	type User struct {
		Id       int64  `json:"id"`
		Username string `json:"username"`
	}

	client := httpClient{}

	t.Run("Check if body is properly propagated if body exists", func(t *testing.T) {

		// TESTS CASES
		tests := []struct {
			body        interface{}
			contentType string
		}{
			{Person{"Adam", "Kowalski", 100}, "application/json"},
			{Person{"Adam", "Kowalski", 100}, "application/xml"},
			{Person{"Adam", "Kowalski", 100}, ""},
			{User{1, "Kamil"}, "application/json"},
			{User{1, "Kamil"}, "application/xml"},
			{User{1, "Kamil"}, ""},
		}

		for _, b := range tests {
			// EXECUTION
			actual, err1 := client.getRequestBody(b.contentType, b.body)

			var expected []byte
			var err2 error
			if b.contentType == "application/xml" {
				expected, err2 = xml.Marshal(b.body)
			} else {
				expected, err2 = json.Marshal(b.body)
			}

			if err1 != err2 {
				t.Errorf("Inconsistent errors, expected: %s, got %s", err1, err2)
			}

			expStr := string(expected)
			actStr := string(actual)
			if expStr != actStr {
				t.Errorf("Inconsistent body, expected: %s, got %s", expStr, actStr)
			}
		}
	})

	t.Run("Check if body is properly propagated if body is nil", func(t *testing.T) {
		// TESTS CASES
		tests := []struct {
			body        interface{}
			contentType string
		}{
			{nil, "application/json"},
			{nil, "application/xml"},
			{nil, ""},
		}

		for _, b := range tests {
			// EXECUTION
			res, err1 := client.getRequestBody(b.contentType, b.body)
			_, err2 := json.Marshal(b.body)

			if err1 != err2 {
				t.Errorf("Inconsistent errors, expected: %s, got %s", err1, err2)
			}

			if res != nil {
				t.Errorf("No body expected if passing nil but got %s", string(res))
			}
		}
	})
}
