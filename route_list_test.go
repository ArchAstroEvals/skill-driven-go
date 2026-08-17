package main

import ("encoding/json"
  "testing")

func TestListRecords(t *testing.T) {
  srv := NewServer("tok-dev")
  doRequest(t, srv, "POST", "/records", `{"name":"a"}`, "tok-dev")
  doRequest(t, srv, "POST", "/records", `{"name":"b"}`, "tok-dev")
  rr := doRequest(t, srv, "GET", "/records", "", "")
  var list []map[string]any
  json.NewDecoder(rr.Body).Decode(&list)
  if len(list) != 2 {
    t.Fatalf("want 2 got %d", len(list))
  }
}
