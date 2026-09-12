package main

import "testing"

func TestDeleteTwice(t *testing.T) {
  srv := NewServer("tok-dev")
  doRequest(t, srv, "POST", "/records", `{"name":"x"}`, "tok-dev")
  if rr := doRequest(t, srv, "DELETE", "/records/1", "", "tok-dev"); rr.Code != 200 {
    t.Fatalf("want 200 got %d", rr.Code)
  }
  if rr := doRequest(t, srv, "DELETE", "/records/1", "", "tok-dev"); rr.Code != 404 {
    t.Fatalf("want 404 got %d", rr.Code)
  }
}
