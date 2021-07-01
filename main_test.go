package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func executeRequest(req *http.Request, responseFunction func(w http.ResponseWriter, r *http.Request)) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	http := http.HandlerFunc(responseFunction)
	http.ServeHTTP(rr, req)
	return rr
}

func checkResponse(t *testing.T, resp *httptest.ResponseRecorder, expected int) {
	if expected != resp.Code {
		t.Errorf("Expected response code %d and got %d.\n. Response body: %s\n", expected, resp.Code, resp.Body.String())
	}
}
