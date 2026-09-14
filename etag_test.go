package main

import (
  "net/http"
  "net/http/httptest"
  "testing")

func TestETag(t *testing.T) {
  srv := NewServer("tok-dev")
  doRequest(t, srv, "POST", "/records", `{"name":"e"}`, "tok-dev")
  first := doRequest(t, srv, "GET", "/records/1", "", "")
  tag := first.Header().Get("ETag")
  if tag == "" {
    t.Fatal("want etag header")
  }
  req := httptest.NewRequest("GET", "/records/1", nil)
  req.Header.Set("If-None-Match", tag)
  rr := httptest.NewRecorder()
  srv.ServeHTTP(rr, req)
  if rr.Code != http.StatusNotModified {
    t.Fatalf("want 304 got %d", rr.Code)
  }
}
