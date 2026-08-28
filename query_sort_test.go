package main

import ("encoding/json"
  "net/http/httptest"
  "testing")

func decodeList(t *testing.T, rr *httptest.ResponseRecorder, out any) {
  t.Helper()
  if err := json.NewDecoder(rr.Body).Decode(out); err != nil {
    t.Fatal(err)
  }
}

func TestSortWired(t *testing.T) {
  srv := NewServer("tok-dev")
  doRequest(t, srv, "POST", "/records", `{"name":"b"}`, "tok-dev")
  doRequest(t, srv, "POST", "/records", `{"name":"a"}`, "tok-dev")
  rr := doRequest(t, srv, "GET", "/records?sort=name", "", "")
  var list []map[string]any
  decodeList(t, rr, &list)
  if list[0]["name"] != "a" || list[1]["name"] != "b" {
    t.Fatalf("want sorted got %v", list)
  }
}
