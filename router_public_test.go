package main

import "testing"

func TestReadsPublic(t *testing.T) {
  srv := NewServer("tok-dev")
  if rr := doRequest(t, srv, "GET", "/records", "", ""); rr.Code != 200 {
    t.Fatalf("want 200 got %d", rr.Code)
  }
  if rr := doRequest(t, srv, "GET", "/health", "", ""); rr.Code != 200 {
    t.Fatalf("want 200 got %d", rr.Code)
  }
}
