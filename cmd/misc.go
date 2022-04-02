package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/Cyb3r-Jak3/common/v4"
)

//GravatarRequestBody represents the expected incoming JSON
type GravatarRequestBody struct {
	Email string `json:"email"`
}

//GravatarResponseBody represents the outgoing JSON
type GravatarResponseBody struct {
	Hash string `json:"hash"`
}

func miscGravatarHash(w http.ResponseWriter, req *http.Request) {
	req.Body = http.MaxBytesReader(w, req.Body, 1*1024*1024)
	if req.Body == http.NoBody || req.ContentLength == 0 {
		http.Error(w, "JSON body required", http.StatusBadRequest)
		return
	}
	out, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var RequestBody GravatarRequestBody
	if err := json.Unmarshal(out, &RequestBody); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	hash := md5.Sum([]byte(strings.TrimSpace(RequestBody.Email)))

	common.JSONMarshalResponse(w, &GravatarResponseBody{Hash: hex.EncodeToString(hash[:])})
}

//StringRequestBody represents the expected incoming JSON
type StringRequestBody struct {
	String       string `json:"string"`
	Modification string `json:"modification"`
}

//StringResponseBody represents the outgoing JSON
type StringResponseBody struct {
	String string `json:"string"`
}

func miscStringChange(w http.ResponseWriter, req *http.Request) {
	req.Body = http.MaxBytesReader(w, req.Body, 1*1024*1024)
	if req.Body == http.NoBody || req.ContentLength == 0 {
		http.Error(w, "JSON body required", http.StatusBadRequest)
		return
	}
	out, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var RequestBody StringRequestBody
	if err := json.Unmarshal(out, &RequestBody); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	switch RequestBody.Modification {
	case "lower", "l":
		common.JSONMarshalResponse(w, &StringResponseBody{strings.ToLower(RequestBody.String)})
	case "upper", "u":
		common.JSONMarshalResponse(w, &StringResponseBody{strings.ToUpper(RequestBody.String)})
	case "title", "t":
		common.JSONMarshalResponse(w, &StringResponseBody{strings.Title(RequestBody.String)})
	default:
		http.Error(w, "Need to specify a modification", http.StatusBadRequest)
	}
}

func VersionInfo(w http.ResponseWriter, _ *http.Request) {
	common.JSONResponse(w, []byte(`{"Version": "`+Version+`", "Date": "`+Date+`", "Commit": "`+Commit+`"}`))
}
