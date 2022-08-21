package go_github_api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func fetch(url string) ([]byte, error) {
	client := http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "go_github_api")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.Body != nil {
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		return body, nil
	}
	return nil, nil
}

func GetPublicRepositories(user string) ([]*Repository, error) {
	var repos []*Repository
	url := fmt.Sprintf("https://api.github.com/users/%s/repos", user)

	res, err := fetch(url)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(res, &repos); err != nil {
		return nil, err
	}
	return repos, nil
}
