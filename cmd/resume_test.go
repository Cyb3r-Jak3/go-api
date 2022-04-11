package main

import (
	"bytes"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"testing"

	"github.com/Cyb3r-Jak3/common/v4"
)

func TestGetResume(t *testing.T) {
	// Get empty resume
	getResume()
	// Use already downloaded resume
	getResume()
}

func TestEncryptResume(t *testing.T) {
	_, err := common.DownloadFile("https://cyberjake.xyz/keys/WebsitePublic.asc", "key.asc")
	if err != nil {
		t.Errorf("Error downloading resume: %s", err)
	}
	file, _ := os.Open("key.asc")
	fileContents, _ := ioutil.ReadAll(file)
	if err = file.Close(); err != nil {
		t.Errorf("Error closing downloaded resume: %s", err)
	}
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("key", "key.asc")
	if _, err = part.Write(fileContents); err != nil {
		t.Errorf("Error writing file to request: %s", err)
	}
	r, _ := http.NewRequest("POST", "/", body)
	r.Header.Add("Content-Type", writer.FormDataContentType())
	if err = writer.Close(); err != nil {
		t.Errorf("Closing multipart writer")
	}
	rr := executeRequest(r, encryptResume)
	checkResponse(t, rr, http.StatusOK)
}

func BenchmarkEncryptResume(b *testing.B) {
	_, _ = common.DownloadFile("https://www.jwhite.network/keys/WebsitePublic.asc", "key.asc")
	file, _ := os.Open("key.asc")
	fileContents, _ := ioutil.ReadAll(file)
	_ = file.Close()
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("key", "key.asc")
	_, _ = part.Write(fileContents)
	r, _ := http.NewRequest("POST", "/", body)
	r.Header.Add("Content-Type", writer.FormDataContentType())
	_ = writer.Close()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		executeRequest(r, encryptResume)
	}

}

func TestEncryptResumeBadFile(t *testing.T) {
	file, _ := os.Open("main.go")
	fileContents, _ := ioutil.ReadAll(file)
	if err := file.Close(); err != nil {
		t.Errorf("Error closing downloaded resume: %s", err)
	}
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("key", "main")
	if _, err := part.Write(fileContents); err != nil {
		t.Errorf("Error writing file to request: %s", err)
	}
	r, _ := http.NewRequest("POST", "/", body)
	r.Header.Add("Content-Type", writer.FormDataContentType())
	if err := writer.Close(); err != nil {
		t.Errorf("Closing multipart writer")
	}
	rr := executeRequest(r, encryptResume)
	checkResponse(t, rr, http.StatusInternalServerError)
}

func TestEncryptResumeEmpty(t *testing.T) {
	r, _ := http.NewRequest("POST", "/", nil)
	rr := executeRequest(r, encryptResume)
	checkResponse(t, rr, http.StatusBadRequest)
}
