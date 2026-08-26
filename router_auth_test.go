package main

import "testing"

func TestWritesNeedTokens(t *testing.T) {
  srv := NewServer("tok-dev")
  if rr := doRequest(t, srv, "POST", "/records", `{"name":"x"}`, "bad"); rr.Code != 401 {
    t.Fatalf("want 401 got %d", rr.Code)
  }
  if rr := doRequest(t, srv, "POST", "/records", `{"name":"x"}`, ""); rr.Code != 401 {
    t.Fatalf("want 401 got %d", rr.Code)
  }
  if rr := doRequest(t, srv, "DELETE", "/records/1", "", ""); rr.Code != 401 {
    t.Fatalf("want 401 got %d", rr.Code)
  }
  if rr := doRequest(t, srv, "POST", "/records", `{"name":"x"}`, "tok-dev"); rr.Code != 201 {
    t.Fatalf("want 201 got %d", rr.Code)
  }
}
