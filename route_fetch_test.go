package main

import "testing"

func TestFetchRecord(t *testing.T) {
  srv := NewServer("tok-dev")
  doRequest(t, srv, "POST", "/records", `{"name":"Widget"}`, "tok-dev")
  rr := doRequest(t, srv, "GET", "/records/1", "", "")
  if decode(t, rr)["name"] != "Widget" {
    t.Fatal("want widget")
  }
  bad := doRequest(t, srv, "GET", "/records/99", "", "")
  if bad.Code != 404 {
    t.Fatal("want 404")
  }
  malformed := doRequest(t, srv, "GET", "/records/abc", "", "")
  if malformed.Code != 404 {
    t.Fatal("want 404 for abc")
  }
}
