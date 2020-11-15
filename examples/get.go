package examples

import (
	"fmt"
)

type Endpoints struct {
	CurrentUserUrl    string `json:"current_user_url"`
	AuthorizationsUrl string `json:"authorizations_url"`
	RepositoryUrl     string `json:"repository_url"`
}

func GetEndpoints() (*Endpoints, error) {
	res, err := httpClient.Get("https://api.github.com")
	if err != nil {
		// Deal with the error as you need
		return nil, err
	}
	fmt.Println(fmt.Sprintf("Status Code: %d\n", res.StatusCode))
	fmt.Println(fmt.Sprintf("Status: %s\n", res.Status))
	fmt.Println(fmt.Sprintf("Body: %s\n", res.String()))

	var endpoints Endpoints
	if err := res.UnmarshalJSON(&endpoints); err != nil {
		// Deal with unmarshal error as you need
		return nil, err
	}
	fmt.Println(fmt.Sprintf("Repository URL: %s", endpoints.RepositoryUrl))
	return &endpoints, err
}
