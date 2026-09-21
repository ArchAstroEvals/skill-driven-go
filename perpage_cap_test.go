package main

import "testing"

func TestPerPageCap(t *testing.T) {
  srv := NewServer("tok-dev")
  doRequest(t, srv, "POST", "/records", `{"name":"a"}`, "tok-dev")
  doRequest(t, srv, "POST", "/records", `{"name":"b"}`, "tok-dev")
  rr := doRequest(t, srv, "GET", "/records?per_page=500", "", "")
  var list []map[string]any
  decodeList(t, rr, &list)
  if len(list) != 2 {
    t.Fatalf("want both got %v", list)
  }
}
