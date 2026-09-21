package main

import "testing"

func TestSearchRoute(t *testing.T) {
  srv := NewServer("tok-dev")
  doRequest(t, srv, "POST", "/records", `{"name":"Widget"}`, "tok-dev")
  doRequest(t, srv, "POST", "/records", `{"name":"Gadget"}`, "tok-dev")
  rr := doRequest(t, srv, "GET", "/records?q=wid", "", "")
  var list []map[string]any
  decodeList(t, rr, &list)
  if len(list) != 1 || list[0]["name"] != "Widget" {
    t.Fatalf("want widget got %v", list)
  }
}
