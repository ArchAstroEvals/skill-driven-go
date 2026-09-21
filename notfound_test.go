package main

import "testing"

func TestNotFoundShape(t *testing.T) {
  srv := NewServer("tok-dev")
  rr := doRequest(t, srv, "GET", "/records/99", "", "")
  if rr.Code != 404 {
    t.Fatalf("want 404 got %d", rr.Code)
  }
  body := decode(t, rr)
  if body["error"] != "not_found" || body["resource"] != "record" {
    t.Fatalf("want shaped 404 got %v", body)
  }
}
