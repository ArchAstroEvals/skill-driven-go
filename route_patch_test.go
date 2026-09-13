package main

import "testing"

func TestPatch(t *testing.T) {
  srv := NewServer("tok-dev")
  created := decode(t, doRequest(t, srv, "POST", "/records", `{"name":"old"}`, "tok-dev"))
  rr := doRequest(t, srv, "PATCH", "/records/1", `{"name":"new"}`, "tok-dev")
  if rr.Code != 200 || decode(t, rr)["name"] != "new" {
    t.Fatalf("want rename got %d", rr.Code)
  }
  if created["name"] != "old" {
    t.Fatal("want original kept")
  }
  if rr := doRequest(t, srv, "PATCH", "/records/99", `{"name":"x"}`, "tok-dev"); rr.Code != 404 {
    t.Fatalf("want 404 got %d", rr.Code)
  }
  if rr := doRequest(t, srv, "PATCH", "/records/1", `{"name":"new"}`, ""); rr.Code != 401 {
    t.Fatalf("want 401 got %d", rr.Code)
  }
}
