package main

import (
  "encoding/json"
  "net/http/httptest"
  "strings"
  "testing"
)

func decode(t *testing.T, rr *httptest.ResponseRecorder) map[string]any {
  t.Helper()
  var m map[string]any
  if err := json.NewDecoder(rr.Body).Decode(&m); err != nil {
    t.Fatal(err)
  }
  return m
}

func doRequest(t *testing.T, srv *Server, method, path, body, token string) *httptest.ResponseRecorder {
  t.Helper()
  var reader *strings.Reader
  if body == "" {
    reader = strings.NewReader("")
  } else {
    reader = strings.NewReader(body)
  }
  req := httptest.NewRequest(method, path, reader)
  if token != "" {
    req.Header.Set("Authorization", "Bearer "+token)
  }
  rr := httptest.NewRecorder()
  srv.ServeHTTP(rr, req)
  return rr
}

func TestUnknownPathIs404(t *testing.T) {
  srv := NewServer("tok-dev")
  rr := doRequest(t, srv, "GET", "/nope", "", "")
  if rr.Code != 404 {
    t.Fatalf("want 404 got %d", rr.Code)
  }
}
