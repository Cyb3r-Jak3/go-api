package main

import (
	"io"
	"net/http"
	"os"

	"github.com/Cyb3r-Jak3/common/v4"
	"github.com/ProtonMail/gopenpgp/v2/helper"
)

const (
	resumeURL      = "https://www.jwhite.network/resumes/JacobWhiteResume.pdf"
	resumeFilePath = "./resume.pdf"
)

func getResume() {
	if _, err := os.Stat(resumeFilePath); !os.IsNotExist(err) {
		log.Debug("Don't need to download resume")
		return
	}
	log.Debug("Downloading resume")
	ok, err := common.DownloadFile(resumeURL, resumeFilePath)
	if err != nil {
		log.WithError(err).Fatal("Error download resume")
	}
	if !ok {
		log.Fatal("Download reported failed")
	}
}

func encryptResume(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		redirect(w, req)
		return
	}
	// Limit request body to 5 KB
	req.Body = http.MaxBytesReader(w, req.Body, 5000)
	file, _, err := req.FormFile("key")
	if err != nil {
		httpError(w, err, "Error reading the form body", http.StatusBadRequest)
		return
	}
	defer file.Close()
	fileContent, err := io.ReadAll(file)
	if err != nil {
		httpError(w, err, "Error reading the key body", http.StatusBadRequest)
		return
	}
	resume, err := os.ReadFile(resumeFilePath)
	if err != nil {
		httpError(w, err, "Error reading saved resume", http.StatusBadRequest)
		return
	}
	returned, err := helper.EncryptBinaryMessageArmored(string(fileContent), resume)
	if err != nil {
		httpError(w, err, "Error encrypting resume", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Disposition", "attachment; filename=jwhite_signed_resume.pdf.gpg")
	common.ContentResponse(w, "application/octet-stream", []byte(returned))

}
