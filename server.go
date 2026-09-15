package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	store *Store
	token string
	mux   *http.ServeMux
	handler http.Handler
	bucket *Bucket
}

func NewServer(token string) *Server {
	s := &Server{store: NewStore(), token: token, mux: http.NewServeMux(), bucket: NewBucket(10000, time.Minute)}
	s.mux.HandleFunc("GET /health", s.handleHealth)
	s.mux.HandleFunc("GET /version", s.handleVersion)
	s.mux.HandleFunc("POST /records", s.limit(s.requireAuth(s.handleCreate)))
	s.mux.HandleFunc("GET /records/{id}", s.handleFetch)
	s.mux.HandleFunc("GET /records", s.handleList)
	s.mux.HandleFunc("DELETE /records/{id}", s.requireAuth(s.handleDelete))
	s.mux.HandleFunc("PATCH /records/{id}", s.requireAuth(s.handlePatch))
	s.mux.HandleFunc("OPTIONS /records", s.handleOptions)
	s.mux.HandleFunc("OPTIONS /records/{id}", s.handleOptions)
	s.handler = withLogging(withRequestID(withCORS(withMaxBytes(s.mux, 1<<20))))
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.handler.ServeHTTP(w, r)
}

func (s *Server) handleOptions(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, 204, map[string]any{})
}

func (s *Server) requireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !validToken(r.Header.Get("Authorization"), s.token) {
			writeJSON(w, 401, unauthorized())
			return
		}
		next(w, r)
	}
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
		writeJSON(w, 404, notFound("route"))
		return
	}
	rec, ok := s.store.Get(id)
	if !ok {
		writeJSON(w, 404, notFound("record"))
		return
	}
	etag := "W/\"" + strconv.Itoa(id) + "\""
	w.Header().Set("ETag", etag)
	if r.Header.Get("If-None-Match") == etag {
		w.WriteHeader(304)
		return
	}
	writeJSON(w, 200, rec)
}

func (s *Server) handleList(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	params := map[string]string{}
	for k, v := range q {
		if len(v) > 0 {
			params[k] = v[0]
		}
	}
	recs := filterKnown(s.store.List(), params)
	if sf := params["sort"]; sf != "" {
		recs = sortBy(recs, sf, params["order"] == "desc")
	}
	perPage := atoiOr(params["per_page"], 20)
	if perPage > 100 {
		perPage = 100
	}
	w.Header().Set("X-Total-Count", strconv.Itoa(len(recs)))
	writeJSON(w, 200, paginate(recs, atoiOr(params["page"], 1), perPage))
}

func (s *Server) handleDelete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeJSON(w, 404, notFound("route"))
		return
	}
	rec, ok := s.store.Get(id)
	if !ok {
		writeJSON(w, 404, notFound("record"))
		return
	}
	s.store.Delete(id)
	writeJSON(w, 200, rec)
}
func (s *Server) handlePatch(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeJSON(w, 404, notFound("route"))
		return
	}
	var body map[string]any
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeErr(w, 400, "bad_request")
		return
	}
	if name, present := body["name"]; present {
		if text, ok := name.(string); !ok || text == "" {
			writeErr(w, 400, "invalid_name")
			return
		}
	}
	rec, ok := s.store.Update(id, body)
	if !ok {
		writeJSON(w, 404, notFound("record"))
		return
	}
	writeJSON(w, 200, rec)
}

