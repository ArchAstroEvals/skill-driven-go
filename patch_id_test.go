package main

import "testing"

func TestPatchIdImmutable(t *testing.T) {
  srv := NewServer("tok-dev")
  doRequest(t, srv, "POST", "/records", `{"name":"a"}`, "tok-dev")
  doRequest(t, srv, "PATCH", "/records/1", `{"id":99,"name":"b"}`, "tok-dev")
  if got := decode(t, doRequest(t, srv, "GET", "/records/1", "", "")); got["id"] != float64(1) {
    t.Fatalf("want id 1 got %v", got)
  }
}
