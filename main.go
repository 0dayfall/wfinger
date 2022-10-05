package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

var data map[string]JRD

func main() {
	setup()
	log.Fatal(http.ListenAndServeTLS(":443", filepath.Join("cert", "go-server.crt"), filepath.Join("cert", "go-server.key"), http.HandlerFunc(ServeHTTP)))
}

func finger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resource := r.URL.Query().Get("resource")
	bytes, err := json.Marshal(getJRD(getAccountName(resource)))
	if err != nil {
		log.Fatal(err)
	}
	_, err = w.Write(bytes)
	if err != nil {
		log.Fatal(err)
	}
}

func update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var jrd JRD
	err = json.Unmarshal(body, &jrd)
	if err != nil {
		log.Fatal(err)
	}
}
