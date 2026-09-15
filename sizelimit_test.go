package main

import (
  "strings"
  "testing")

func TestSizeLimit(t *testing.T) {
  srv := NewServer("tok-dev")
  big := `{"name":"` + strings.Repeat("x", 2<<20) + `"}`
  if rr := doRequest(t, srv, "POST", "/records", big, "tok-dev"); rr.Code != 400 {
    t.Fatalf("want 400 got %d", rr.Code)
  }
}
