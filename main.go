package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe("", &NewServer(mux.NewRouter()).router))
}