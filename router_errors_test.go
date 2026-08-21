package main

import "testing"

func TestRouterErrorBodies(t *testing.T) {
  srv := NewServer("tok-dev")
  rr := doRequest(t, srv, "GET", "/records/99", "", "")
  if decode(t, rr)["error"] != "not_found" {
    t.Fatal("want error code")
  }
}
