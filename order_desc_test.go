package main

import "testing"

func TestOrderDesc(t *testing.T) {
  srv := NewServer("tok-dev")
  doRequest(t, srv, "POST", "/records", `{"name":"a"}`, "tok-dev")
  doRequest(t, srv, "POST", "/records", `{"name":"b"}`, "tok-dev")
  rr := doRequest(t, srv, "GET", "/records?sort=name&order=desc", "", "")
  var list []map[string]any
  decodeList(t, rr, &list)
  if list[0]["name"] != "b" || list[1]["name"] != "a" {
    t.Fatalf("want desc got %v", list)
  }
}
