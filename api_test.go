package go_github_api

import (
	"testing"
)

func TestGetPublicRepositories(t *testing.T) {
	user := ""
	_, err := GetPublicRepositories(user)
	if err != nil {
		t.Fatalf(err.Error())
	}
}
