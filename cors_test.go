package main

import "testing"

func TestCORS(t *testing.T) {
  srv := NewServer("tok-dev")
  rr := doRequest(t, srv, "GET", "/health", "", "")
  if rr.Header().Get("Access-Control-Allow-Origin") != "*" {
    t.Fatal("want cors header")
  }
}
