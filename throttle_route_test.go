package main

import ("testing"
  "time")

func TestRouteLimit(t *testing.T) {
  srv := NewServer("tok-dev")
  srv.bucket = NewBucket(1, time.Hour)
  first := doRequest(t, srv, "POST", "/records", `{"name":"a"}`, "tok-dev")
  if first.Code != 201 {
    t.Fatalf("want 201 got %d", first.Code)
  }
  second := doRequest(t, srv, "POST", "/records", `{"name":"b"}`, "tok-dev")
  if second.Code != 429 {
    t.Fatalf("want 429 got %d", second.Code)
  }
}
