package main

import (
	"net/http"

	"github.com/Cyb3r-Jak3/common/v4"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
)

var (
	CORSDomains = []string{"https://*.cyberjake.xyz"}
	log         = logrus.New()
	host        string
	port        string
	c           *cors.Cors
	Version     = ""
	Date        = ""
	Commit      = ""
)

func httpError(w http.ResponseWriter, err error, message string, statusCode int) {
	log.WithError(err).Error(message)
	http.Error(w, err.Error(), statusCode)
}

func redirect(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "https://cyberjake.xyz", http.StatusPermanentRedirect)
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
	r.HandleFunc("/version", VersionInfo)
	r.NotFoundHandler = http.HandlerFunc(redirect)
	r.HandleFunc("/encrypted_resume", common.AllowedMethods(encryptResume, "POST,OPTIONS"))
	r.HandleFunc("/misc/string", common.AllowedMethods(miscStringChange, "POST,OPTIONS"))
	bindAddress := host + ":" + port
	log.Info("Starting on " + bindAddress)
	if err := http.ListenAndServe(bindAddress, c.Handler(r)); err != nil {
		log.WithError(err).Fatal("Error running server")
	}

}
