package main

import "net/http"

type Server struct {
	token string
	mux   *http.ServeMux
}

func NewServer(token string) *Server {
	s := &Server{token: token, mux: http.NewServeMux()}
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
