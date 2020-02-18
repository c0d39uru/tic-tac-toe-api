package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	version string
	router  mux.Router
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func NewServer(router *mux.Router) *Server {
	s := &Server{
		version: "1.0",
		router: *router,
	}

	s.router.HandleFunc("/ping", s.handlePing()).Methods("GET")
	s.router.HandleFunc("/play", s.handlePlay()).Methods("POST")

	return s
}
