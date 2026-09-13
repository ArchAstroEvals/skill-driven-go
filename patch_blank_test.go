package main

import "testing"

func TestPatchBlank(t *testing.T) {
  srv := NewServer("tok-dev")
  doRequest(t, srv, "POST", "/records", `{"name":"keep"}`, "tok-dev")
  if rr := doRequest(t, srv, "PATCH", "/records/1", `{"name":""}`, "tok-dev"); rr.Code != 400 {
    t.Fatalf("want 400 got %d", rr.Code)
  }
  rr := doRequest(t, srv, "GET", "/records/1", "", "")
  if decode(t, rr)["name"] != "keep" {
    t.Fatal("want name unchanged")
  }
}
