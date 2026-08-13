package main

import "testing"

func TestHealth(t *testing.T) {
  srv := NewServer("tok-dev")
  rr := doRequest(t, srv, "GET", "/health", "", "")
  if rr.Code != 200 {
    t.Fatalf("want 200 got %d", rr.Code)
  }
  if decode(t, rr)["status"] != "ok" {
    t.Fatal("want status ok")
  }
}
