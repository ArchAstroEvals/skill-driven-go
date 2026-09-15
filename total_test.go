package main

import "testing"

func TestTotalCount(t *testing.T) {
  srv := NewServer("tok-dev")
  doRequest(t, srv, "POST", "/records", `{"name":"a"}`, "tok-dev")
  doRequest(t, srv, "POST", "/records", `{"name":"b"}`, "tok-dev")
  doRequest(t, srv, "POST", "/records", `{"name":"c"}`, "tok-dev")
  rr := doRequest(t, srv, "GET", "/records?per_page=2", "", "")
  if rr.Header().Get("X-Total-Count") != "3" {
    t.Fatalf("want total 3 got %q", rr.Header().Get("X-Total-Count"))
  }
}
