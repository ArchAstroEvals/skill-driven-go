package main

import "testing"

func TestMethodMismatch(t *testing.T) {
  srv := NewServer("tok-dev")
  if rr := doRequest(t, srv, "POST", "/health", "", ""); rr.Code != 405 {
    t.Fatalf("want 405 got %d", rr.Code)
  }
  if rr := doRequest(t, srv, "PUT", "/records", "", ""); rr.Code != 405 {
    t.Fatalf("want 405 got %d", rr.Code)
  }
}
