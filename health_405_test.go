package main

import "testing"

func TestHealth405(t *testing.T) {
  srv := NewServer("tok-dev")
  rr := doRequest(t, srv, "POST", "/health", "", "")
  if rr.Code != 405 {
    t.Fatalf("want 405 got %d", rr.Code)
  }
}
