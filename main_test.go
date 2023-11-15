package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEchoHandler(t *testing.T) {
	tt := []struct {
		echoRequest  string
		echoResponse string
	}{
		{"hello", "hello"},
		{"hi", "hi"},
		{"kangho", "kangho"},
	}

	for _, tc := range tt {
		path := fmt.Sprintf("/%s", tc.echoRequest)
		req, err := http.NewRequest("GET", path, nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()

		router := mux.NewRouter()
		router.HandleFunc("/{echo}", EchoHandler)
		router.ServeHTTP(rr, req)
		var result struct {
			FromServer string `json:"from_server"`
		}
		if err := json.NewDecoder(rr.Body).Decode(&result); err != nil {
			log.Fatalln(err)
		}
		if rr.Code != http.StatusOK || tc.echoResponse != result.FromServer {
			t.Errorf("Error !")
		}
	}
}
