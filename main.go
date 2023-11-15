package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"from_server": "%v"}`, vars["echo"])
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{echo}", EchoHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8008", r))
}
