package main

import "testing"

func TestHealthJSON(t *testing.T) {
  srv := NewServer("tok-dev")
  rr := doRequest(t, srv, "GET", "/health", "", "")
  if ct := rr.Header().Get("Content-Type"); ct != "application/json" {
    t.Fatalf("want json got %q", ct)
  }
}
