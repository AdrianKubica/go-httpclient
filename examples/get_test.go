package examples

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/AdrianKubica/go-httpclient/gohttp_mock"
)

func TestMain(m *testing.M) {
	// Tell the HTTP library to mock any further request from here before each test case
	gohttp_mock.StartMockServer()

	os.Exit(m.Run())
}

func TestGetEndpoints(t *testing.T) {
	t.Run("Error fetching from github", func(t *testing.T) {
		// Initialization
		gohttp_mock.DeleteMocks()
		gohttp_mock.AddMock(gohttp_mock.Mock{
			Method: http.MethodGet,
			Url:    "https://api.github.com",
			Error:  errors.New("timeout getting github endpoint"),
		})

		// Execution
		endpoints, err := GetEndpoints()

		// Validation
		if endpoints != nil {
			t.Error("no endpoints expected at this point")
		}

		if err == nil {
			t.Error("an error was expected")
		}

		if err.Error() != "timeout getting github endpoint" {
			t.Error("invalid error message received")
		}
	})

	t.Run("Error unmarshal response body", func(t *testing.T) {
		// Initialization
		gohttp_mock.DeleteMocks()
		gohttp_mock.AddMock(gohttp_mock.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"current_user_url": 123}`,
		})
		// Execution
		endpoints, err := GetEndpoints()

		// Validation
		if endpoints != nil {
			t.Error("no endpoints expected at this point")
		}

		if err == nil {
			t.Error("an error was expected")
		}

		if !strings.Contains(err.Error(), "cannot unmarshal number into Go struct field") {
			t.Error("invalid error message received")
		}
	})

	t.Run("No error", func(t *testing.T) {
		// Initialization
		gohttp_mock.DeleteMocks()
		gohttp_mock.AddMock(gohttp_mock.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"current_user_url": "https://api.github.com/user"}`,
		})
		// Execution
		endpoints, err := GetEndpoints()

		// Validation
		if err != nil {
			t.Error(fmt.Sprintf("no error was expected and we got '%s'", err.Error()))
		}

		if endpoints == nil {
			t.Error("endpoints were expected and we got nil")
		}

		if endpoints.CurrentUserUrl != "https://api.github.com/user" {
			t.Error("invalid current user url")
		}
	})
}
