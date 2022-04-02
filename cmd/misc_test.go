package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/Cyb3r-Jak3/common/v4"
)

func TestMiscGravatar(t *testing.T) {
	r, _ := http.NewRequest("POST", "/", bytes.NewBuffer([]byte(`{"email": "git@cyberjake.xyz"}`)))
	r.Header.Set("Content-Type", "application/json")
	rr := executeRequest(r, miscGravatarHash)
	checkResponse(t, rr, http.StatusOK)
	var bodyResponse GravatarResponseBody
	err := json.Unmarshal(rr.Body.Bytes(), &bodyResponse)
	if err != nil {
		t.Errorf("Unable to marshal response JSON. %s\n", err)
	}
	correctHash := "53bb2d43885821c16259c5311d3755d8"
	if bodyResponse.Hash != correctHash {
		t.Errorf("Got wrong Gravatar hash. Wanted %s Got %s", correctHash, bodyResponse.Hash)
	}

}

func BenchmarkMiscGravatar(b *testing.B) {
	r, _ := http.NewRequest("POST", "/", bytes.NewBuffer([]byte(`{"email": "git@cyberjake.xyz"}`)))
	r.Header.Set("Content-Type", "application/json")
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		executeRequest(r, miscGravatarHash)
	}
}

type TestString struct {
	String         string
	Modification   string
	ExpectedResult string
}

var stringTests = []TestString{
	{
		String:         "HELLO WORLD",
		Modification:   "lower",
		ExpectedResult: "hello world",
	},
	{
		String:         "hello world",
		Modification:   "upper",
		ExpectedResult: "HELLO WORLD",
	},
	{
		String:         "hello world",
		Modification:   "title",
		ExpectedResult: "Hello World",
	},
}

func TestMiscString(t *testing.T) {
	for _, x := range stringTests {
		r, _ := http.NewRequest("POST", "/", bytes.NewBuffer([]byte(fmt.Sprintf(`{"string": "%s", "modification": "%s"}`, x.String, x.Modification))))
		r.Header.Set("Content-Type", "application/json")
		rr := executeRequest(r, miscStringChange)
		checkResponse(t, rr, http.StatusOK)
		var response StringResponseBody
		err := json.Unmarshal(rr.Body.Bytes(), &response)
		if err != nil {
			t.Errorf("Unable to marshal response JSON. %s\n", err)
		}
		if response.String != x.ExpectedResult {
			t.Errorf("Error with %s. Wanted %s Got %s", x.Modification, x.ExpectedResult, response.String)
		}
	}
}

func TestVersionInfo(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	rr := executeRequest(r, VersionInfo)
	checkResponse(t, rr, http.StatusOK)
	headers := rr.Header()
	jsonHeader := headers.Get("Content-Type")
	if jsonHeader != common.JSONApplicationType {
		t.Errorf("Version response was not type content. It was %s", jsonHeader)
	}
}

func BenchmarkMiscStringToLower(b *testing.B) {
	r, _ := http.NewRequest("POST", "/", bytes.NewBuffer([]byte(`{"string": "HELLO WORLD", "modification": "l"}`)))
	r.Header.Set("Content-Type", "application/json")
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		executeRequest(r, miscStringChange)
	}
}

func BenchmarkMiscStringToUpper(b *testing.B) {
	r, _ := http.NewRequest("POST", "/", bytes.NewBuffer([]byte(`{"string": "hello world", "modification": "u"}`)))
	r.Header.Set("Content-Type", "application/json")
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		executeRequest(r, miscStringChange)
	}
}

func BenchmarkMiscStringToTitle(b *testing.B) {
	r, _ := http.NewRequest("POST", "/", bytes.NewBuffer([]byte(`{"string": "Hello World", "modification": "t"}`)))
	r.Header.Set("Content-Type", "application/json")
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		executeRequest(r, miscStringChange)
	}
}

func BenchmarkVersionInfo(b *testing.B) {
	r, _ := http.NewRequest("GET", "/", nil)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		executeRequest(r, VersionInfo)
	}
}
