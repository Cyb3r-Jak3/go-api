package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func executeRequest(req *http.Request, responseFunction func(w http.ResponseWriter, r *http.Request)) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	dummyHTTP := http.HandlerFunc(responseFunction)
	dummyHTTP.ServeHTTP(rr, req)
	return rr
}

func checkResponse(t *testing.T, resp *httptest.ResponseRecorder, expected int) {
	if expected != resp.Code {
		t.Errorf("Expected response code %d and got %d.\n. Response body: %s\n", expected, resp.Code, resp.Body.String())
	}
}

func TestRedirect(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	rr := executeRequest(r, redirect)
	checkResponse(t, rr, http.StatusPermanentRedirect)
}

func Test404(t *testing.T) {
	r, _ := http.NewRequest("GET", "/missing", nil)
	rr := executeRequest(r, redirect)
	checkResponse(t, rr, http.StatusPermanentRedirect)
}
