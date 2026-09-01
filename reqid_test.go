package main

import "testing"

func TestRequestID(t *testing.T) {
  srv := NewServer("tok-dev")
  rr := doRequest(t, srv, "GET", "/health", "", "")
  if rr.Header().Get("X-Request-ID") == "" {
    t.Fatal("want request id header")
  }
  rr2 := doRequest(t, srv, "GET", "/health", "", "")
  rr2.Header().Set("X-Request-ID", "keep-me")
}
