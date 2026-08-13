package main

import "net/http"

type Server struct {
	store *Store
	token string
	mux   *http.ServeMux
}

func NewServer(token string) *Server {
	s := &Server{store: NewStore(), token: token, mux: http.NewServeMux()}
	s.mux.HandleFunc("GET /health", s.handleHealth)
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, 200, map[string]any{"status": "ok"})
}
