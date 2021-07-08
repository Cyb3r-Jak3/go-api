package main

import (
	"context"
	"net/http"
	"time"

	common "github.com/Cyb3r-Jak3/common/go"
	"github.com/google/go-github/v35/github"
)

type condensedRepoInfo struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type reposCache struct {
	repos             []*github.Repository
	CheckTime         time.Time
	condensedRepoInfo []condensedRepoInfo
}

type extendedUserInfo struct {
	*github.User
	GitLabURL string `json:"gitlab_url"`
	CheckTime time.Time
}

var (
	githubClient  = github.NewClient(&http.Client{})
	gitUserInfo   = &extendedUserInfo{}
	gitReposCache = &reposCache{}
)

const (
	githubUser = "Cyb3r-Jak3"
)

func gitRepos(w http.ResponseWriter, req *http.Request) {
	now := time.Now()
	cacheTime := now.Add(-1 * time.Hour)
	if gitReposCache.CheckTime.After(cacheTime) {
		log.Debug("Serveing cache repos")
		common.JSONMarshalResponse(w, gitReposCache.repos)
		return
	}
	repos, _, err := githubClient.Repositories.List(context.TODO(), githubUser, nil)
	if err != nil {
		httpError(w, err, "Error getting repos", http.StatusInternalServerError)
		return
	}
	gitReposCache = &reposCache{
		repos:     repos,
		CheckTime: now,
	}
	common.JSONMarshalResponse(w, gitReposCache.repos)
}

func gitReposList(w http.ResponseWriter, req *http.Request) {
	now := time.Now()
	cacheTime := now.Add(-1 * time.Hour)
	if gitReposCache.CheckTime.After(cacheTime) {
		log.Debug("Serveing cache repos list")
		common.JSONMarshalResponse(w, gitReposCache.condensedRepoInfo)
		return
	}
	repos, _, err := githubClient.Repositories.List(context.TODO(), githubUser, nil)
	if err != nil {
		httpError(w, err, "Error getting repos", http.StatusInternalServerError)
		return
	}
	var responseBody []condensedRepoInfo
	for _, i := range repos {
		responseBody = append(responseBody, condensedRepoInfo{Name: i.GetName(), URL: i.GetURL()})
	}
	gitReposCache = &reposCache{
		repos:             repos,
		CheckTime:         now,
		condensedRepoInfo: responseBody,
	}
	common.JSONMarshalResponse(w, gitReposCache.condensedRepoInfo)

}
func gitUser(w http.ResponseWriter, req *http.Request) {
	now := time.Now()
	cacheTime := now.Add(-1 * time.Hour)
	if gitUserInfo.CheckTime.After(cacheTime) {
		log.Debug("Serveing cache user")
		common.JSONMarshalResponse(w, gitUserInfo)
		return
	}
	user, _, err := githubClient.Users.Get(context.TODO(), githubUser)
	if err != nil {
		httpError(w, err, "Error gettings GitHub User", http.StatusInternalServerError)
		return
	}
	publicEmail := "cyb3rjak3@pm.me"
	user.Email = &publicEmail
	user.URL = user.HTMLURL
	gitUserInfo = &extendedUserInfo{
		User:      user,
		GitLabURL: "https://gitlab.com/Cyb3r-Jak3",
		CheckTime: now,
	}
	common.JSONMarshalResponse(w, gitUserInfo)

}
