package main

import "testing"

func TestUnauthorizedShape(t *testing.T) {
  srv := NewServer("tok-dev")
  rr := doRequest(t, srv, "POST", "/records", `{"name":"x"}`, "")
  if rr.Code != 401 {
    t.Fatalf("want 401 got %d", rr.Code)
  }
  if decode(t, rr)["error"] != "unauthorized" {
    t.Fatal("want unauthorized code")
  }
}
