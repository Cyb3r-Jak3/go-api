package main

import (
	"context"
	"net/http"

	common "github.com/Cyb3r-Jak3/common/go"
	"github.com/google/go-github/v35/github"
)

type condensedRepoInfo struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type extendedUserInfo struct {
	*github.User
	GitLabURL string `json:"gitlab_url"`
}

var githubClient = github.NewClient(&http.Client{})

const (
	githubUser = "Cyb3r-Jak3"
)

func gitRepos(w http.ResponseWriter, req *http.Request) {
	repos, _, err := githubClient.Repositories.List(context.TODO(), githubUser, nil)
	if err != nil {
		httpError(w, err, "Error getting repos", http.StatusInternalServerError)
		return
	}
	common.JSONMarshalResponse(w, repos)
}

func gitReposList(w http.ResponseWriter, req *http.Request) {
	repos, _, err := githubClient.Repositories.List(context.TODO(), githubUser, nil)
	if err != nil {
		httpError(w, err, "Error getting repos", http.StatusInternalServerError)
		return
	}
	var responseBody []condensedRepoInfo
	for _, i := range repos {
		responseBody = append(responseBody, condensedRepoInfo{Name: i.GetName(), URL: i.GetURL()})
	}
	common.JSONMarshalResponse(w, responseBody)

}
func gitUser(w http.ResponseWriter, req *http.Request) {
	user, _, err := githubClient.Users.Get(context.TODO(), githubUser)
	if err != nil {
		httpError(w, err, "Error gettings GitHub User", http.StatusInternalServerError)
		return
	}
	publicEmail := "cyb3rjak3@pm.me"
	user.Email = &publicEmail
	user.URL = user.HTMLURL
	extendedInfo := &extendedUserInfo{
		User: user,
	}
	extendedInfo.GitLabURL = "https://gitlab.com/Cyb3r-Jak3"
	common.JSONMarshalResponse(w, extendedInfo)

}
