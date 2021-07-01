package main

import (
	"net/http"
	"testing"

	common "github.com/Cyb3r-Jak3/common/go"
)

func TestGitUser(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	rr := executeRequest(r, gitUser)
	checkResponse(t, rr, http.StatusOK)
	resp := rr.Result()
	if resp.Header.Get("Content-Type") != common.JSONApplicationType {
		t.Errorf("Wanted JSON response and got %s", resp.Header.Get("Content-Type"))
	}
}

func TestGitRepos(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	rr := executeRequest(r, gitRepos)
	checkResponse(t, rr, http.StatusOK)
}

func TestGitReposList(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	rr := executeRequest(r, gitReposList)
	checkResponse(t, rr, http.StatusOK)
}
