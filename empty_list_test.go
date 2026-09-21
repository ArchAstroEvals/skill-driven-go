package main

import "testing"

func TestEmptyList(t *testing.T) {
  srv := NewServer("tok-dev")
  rr := doRequest(t, srv, "GET", "/records", "", "")
  var list []map[string]any
  decodeList(t, rr, &list)
  if list == nil || len(list) != 0 {
    t.Fatalf("want empty array got %v", list)
  }
}
