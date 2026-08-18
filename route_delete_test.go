package main

import "testing"

func TestDeleteRecord(t *testing.T) {
  srv := NewServer("tok-dev")
  doRequest(t, srv, "POST", "/records", `{"name":"x"}`, "tok-dev")
  rr := doRequest(t, srv, "DELETE", "/records/1", "", "tok-dev")
  if rr.Code != 200 {
    t.Fatalf("want 200 got %d", rr.Code)
  }
  again := doRequest(t, srv, "DELETE", "/records/1", "", "tok-dev")
  if again.Code != 404 {
    t.Fatal("want 404 second time")
  }
}
