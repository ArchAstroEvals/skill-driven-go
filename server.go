package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Server struct {
	store *Store
	token string
	mux   *http.ServeMux
}

func NewServer(token string) *Server {
	s := &Server{store: NewStore(), token: token, mux: http.NewServeMux()}
	s.mux.HandleFunc("GET /health", s.handleHealth)
	s.mux.HandleFunc("GET /version", s.handleVersion)
	s.mux.HandleFunc("POST /records", s.handleCreate)
	s.mux.HandleFunc("GET /records/{id}", s.handleFetch)
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, 200, map[string]any{"status": "ok"})
}

func (s *Server) handleVersion(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, 200, map[string]any{"version": "0.1.0"})
}

func (s *Server) handleCreate(w http.ResponseWriter, r *http.Request) {
	var body map[string]any
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeErr(w, 400, "bad_request")
		return
	}
	writeJSON(w, 201, s.store.Create(body))
}

func (s *Server) handleFetch(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeErr(w, 404, "not_found")
		return
	}
	rec, ok := s.store.Get(id)
	if !ok {
		writeErr(w, 404, "not_found")
		return
	}
	writeJSON(w, 200, rec)
}
