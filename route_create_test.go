package main

import "testing"

func TestCreateRecord(t *testing.T) {
  srv := NewServer("tok-dev")
  rr := doRequest(t, srv, "POST", "/records", `{"name":"Widget"}`, "tok-dev")
  if rr.Code != 201 {
    t.Fatalf("want 201 got %d", rr.Code)
  }
  if decode(t, rr)["id"] != float64(1) {
    t.Fatal("want id 1")
  }
}
