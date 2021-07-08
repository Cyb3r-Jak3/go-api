package main

import (
	"bytes"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"testing"

	common "github.com/Cyb3r-Jak3/common/go"
)

func TestGetResume(t *testing.T) {
	// Get empty resume
	getResume()
	// Use already downloaded resume
	getResume()
}

func TestEncryptResume(t *testing.T) {
	common.DownloadFile("https://www.jwhite.network/keys/WebsitePublic.asc", "key.asc")
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
	common.DownloadFile("https://www.jwhite.network/keys/WebsitePublic.asc", "key.asc")
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
