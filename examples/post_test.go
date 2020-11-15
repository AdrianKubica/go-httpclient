package examples

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/AdrianKubica/go-httpclient/gohttp_mock"
)

func TestCreateRepo(t *testing.T) {
	t.Run("error is returned for timeout", func(t *testing.T) {
		gohttp_mock.MockupServer.DeleteMocks()
		gohttp_mock.MockupServer.AddMock(gohttp_mock.Mock{
			Method:      http.MethodPost,
			Url:         "https://api.github.com/user/repos",
			RequestBody: `{"name":"test-repo","private":true}`,

			Error: errors.New("timeout from github"),
		})
		repository := Repository{
			Name:    "test-repo",
			Private: true,
		}
		repo, err := CreateRepo(repository)

		if repo != nil {
			t.Error("no repo expected when we get a timeout from github")
		}

		if err == nil {
			t.Error("an error is expected when we get a timeout from github")
		}

		if err.Error() != "timeout from github" {
			fmt.Println(err.Error())
			t.Error("invalid error message")
		}
	})

	t.Run("repository is properly created", func(t *testing.T) {
		gohttp_mock.MockupServer.DeleteMocks()
		gohttp_mock.MockupServer.AddMock(gohttp_mock.Mock{
			Method:      http.MethodPost,
			Url:         "https://api.github.com/user/repos",
			RequestBody: `{"name":"test-repo","private":true}`,

			ResponseStatusCode: http.StatusCreated,
			ResponseBody:       `{"id":123,"name":"test-repo"}`,
		})
		repository := Repository{
			Name:    "test-repo",
			Private: true,
		}
		repo, err := CreateRepo(repository)

		if err != nil {
			t.Error("no error expected when we get a valid response from github")
		}

		if repo == nil {
			t.Error("a valid repo was expected")
		}

		if repo.Name != repository.Name {
			t.Error("returned invalid repository name from github")
		}
	})
}
