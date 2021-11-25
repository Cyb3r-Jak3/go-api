package main

import (
	"fmt"
	"github.com/Cyb3r-Jak3/common/v4"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"net/http"
)

var (
	CORSDomains = []string{"https://*.jwhite.network", "https://*.cyberjake.xyz"}
	log         = logrus.New()
	host        string
	port        string
	c           *cors.Cors
)

func httpError(w http.ResponseWriter, err error, message string, statusCode int) {
	log.WithError(err).Error(message)
	http.Error(w, err.Error(), statusCode)
}

func redirect(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, fmt.Sprintf("https://%s", req.URL.Host), http.StatusPermanentRedirect)
}

func init() {
	getResume()
	host = common.GetEnv("HOST", "")
	port = common.GetEnv("PORT", "5000")
	prod := common.GetEnv("PRODUCTION", "FALSE") == "TRUE"
	if prod {
		c = cors.New(cors.Options{
			AllowedOrigins: CORSDomains,
		})
		log.SetLevel(logrus.WarnLevel)
	} else {
		c = cors.Default()
		log.SetLevel(logrus.DebugLevel)
	}
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", redirect)
	r.NotFoundHandler = http.HandlerFunc(redirect)
	r.HandleFunc("/encrypted_resume", common.AllowedMethod(encryptResume, "POST,OPTIONS"))
	r.HandleFunc("/git/repos", common.AllowedMethod(gitRepos, "GET,OPTIONS"))
	r.HandleFunc("/git/repos/list", common.AllowedMethod(gitReposList, "GET,OPTIONS"))
	r.HandleFunc("/git/user", common.AllowedMethod(gitUser, "GET,OPTIONS"))
	r.HandleFunc("/misc/gravatar", common.AllowedMethod(miscGravatarHash, "POST,OPTIONS"))
	r.HandleFunc("/misc/string", common.AllowedMethod(miscStringChange, "POST,OPTIONS"))
	log.Info("Starting")
	handler := c.Handler(r)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), handler); err != nil {
		log.WithError(err).Fatal("Error running server")
	}

}
