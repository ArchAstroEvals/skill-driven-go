package main

import "testing"

func TestFilterKind(t *testing.T) {
  srv := NewServer("tok-dev")
  doRequest(t, srv, "POST", "/records", `{"name":"x","kind":"book"}`, "tok-dev")
  doRequest(t, srv, "POST", "/records", `{"name":"y","kind":"movie"}`, "tok-dev")
  rr := doRequest(t, srv, "GET", "/records?kind=book", "", "")
  var list []map[string]any
  decodeList(t, rr, &list)
  if len(list) != 1 || list[0]["name"] != "x" {
    t.Fatalf("want one book got %v", list)
  }
}
