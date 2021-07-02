package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"testing"
)

func TestGetResume(t *testing.T) {
	// Get empty resume
	getResume()
	// Use already downloaded resume
	getResume()
}

func getPublicKey() {
	out, err := os.Create("key.asc")
	if err != nil {
		log.WithError(err).Fatal("Error when creating resume file path")
	}
	resp, err := http.Get("https://www.jwhite.network/keys/WebsitePublic.asc")
	if err != nil {
		log.WithError(err).Fatal("Error downloading resume")
	}
	defer resp.Body.Close()
	if _, err := io.Copy(out, resp.Body); err != nil {
		log.WithError(err).Fatal("Error saving resume")
	}
	if err := out.Close(); err != nil {
		log.WithError(err).Error("Error closing file")
	}
}
func TestEncryptResume(t *testing.T) {
	getPublicKey()
	file, _ := os.Open("key.asc")
	fileContents, _ := ioutil.ReadAll(file)
	file.Close()
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("key", "key.asc")
	part.Write(fileContents)
	r, _ := http.NewRequest("POST", "/", body)
	r.Header.Add("Content-Type", writer.FormDataContentType())
	writer.Close()
	rr := executeRequest(r, encryptResume)
	checkResponse(t, rr, http.StatusOK)
}

func BenchmarkEncryptResume(b *testing.B) {
	getPublicKey()
	file, _ := os.Open("key.asc")
	fileContents, _ := ioutil.ReadAll(file)
	file.Close()
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("key", "key.asc")
	part.Write(fileContents)
	r, _ := http.NewRequest("POST", "/", body)
	r.Header.Add("Content-Type", writer.FormDataContentType())
	writer.Close()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		executeRequest(r, encryptResume)
	}

}

func TestEncryptResumeBadFile(t *testing.T) {
	file, _ := os.Open("main.go")
	fileContents, _ := ioutil.ReadAll(file)
	file.Close()
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("key", "main")
	part.Write(fileContents)
	r, _ := http.NewRequest("POST", "/", body)
	r.Header.Add("Content-Type", writer.FormDataContentType())
	writer.Close()
	rr := executeRequest(r, encryptResume)
	checkResponse(t, rr, http.StatusInternalServerError)
}

func TestEncryptResumeEmpty(t *testing.T) {
	r, _ := http.NewRequest("POST", "/", nil)
	rr := executeRequest(r, encryptResume)
	checkResponse(t, rr, http.StatusBadRequest)
}
