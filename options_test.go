package main

import "testing"

func TestPreflight(t *testing.T) {
  srv := NewServer("tok-dev")
  rr := doRequest(t, srv, "OPTIONS", "/records", "", "")
  if rr.Code != 204 {
    t.Fatalf("want 204 got %d", rr.Code)
  }
}
