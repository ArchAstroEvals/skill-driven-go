package main

import "testing"

func TestVersion(t *testing.T) {
  srv := NewServer("tok-dev")
  rr := doRequest(t, srv, "GET", "/version", "", "")
  if decode(t, rr)["version"] != "0.1.0" {
    t.Fatal("want version 0.1.0")
  }
}
