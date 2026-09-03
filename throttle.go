package main

import (
	"net/http"
	"time"
)

type Bucket struct {
	limit  int
	window time.Duration
	used   int
	reset  time.Time
	empty  bool
}

func NewBucket(limit int, window time.Duration) *Bucket {
	return &Bucket{limit: limit, window: window, empty: true}
}

func (b *Bucket) Allow(now time.Time) bool {
	if b.empty || !now.Before(b.reset) {
		b.empty = false
		b.used = 0
		b.reset = now.Add(b.window)
	}
	if b.used >= b.limit {
		return false
	}
	b.used++
	return true
}

func (s *Server) limit(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !s.bucket.Allow(time.Now()) {
			writeJSON(w, 429, map[string]any{"error": "rate_limited"})
			return
		}
		next(w, r)
	}
}
