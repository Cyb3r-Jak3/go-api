package main

import (
	"fmt"
	"net/http"

	common "github.com/Cyb3r-Jak3/common/go"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var (
	log  = logrus.New()
	host string
	port string
)

func httpError(w http.ResponseWriter, err error, message string, statusCode int) {
	log.WithError(err).Error(message)
	http.Error(w, err.Error(), statusCode)
}

func redirect(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "https://www.jwhite.network", http.StatusPermanentRedirect)
}

func init() {
	getResume()
	host = common.GetEnv("HOST", "")
	port = common.GetEnv("PORT", "5000")
}
func main() {
	log.SetLevel(logrus.DebugLevel)
	r := mux.NewRouter()
	r.HandleFunc("/", redirect)
	r.NotFoundHandler = http.HandlerFunc(redirect)
	r.HandleFunc("/encrypted_resume", common.AllowedMethod(encryptResume, "POST,OPTIONS"))
	r.HandleFunc("/git/repos", common.AllowedMethod(gitRepos, "GET,OPTIONS"))
	r.HandleFunc("/git/repos/list", common.AllowedMethod(gitReposList, "GET,OPTIONS"))
	r.HandleFunc("/git/user", common.AllowedMethod(gitUser, "GET,OPTIONS"))
	log.Info("Starting")
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), r); err != nil {
		log.WithError(err).Fatal("Error running server")
	}

}
