package main

import (
  "encoding/json"
  "fmt"
  "net/http"
  "net/http/httptest"
  "strings"
  "testing")

func intReq(t *testing.T, method, url, body string) *http.Response {
  t.Helper()
  req, err := http.NewRequest(method, url, strings.NewReader(body))
  if err != nil {
    t.Fatal(err)
  }
  req.Header.Set("Authorization", "Bearer tok-int")
  resp, err := http.DefaultClient.Do(req)
  if err != nil {
    t.Fatal(err)
  }
  return resp
}

func TestIntegrationCRUD(t *testing.T) {
  srv := httptest.NewServer(NewServer("tok-int"))
  defer srv.Close()
  resp := intReq(t, "POST", srv.URL+"/records", `{"name":"int"}`)
  if resp.StatusCode != 201 {
    t.Fatalf("want 201 got %d", resp.StatusCode)
  }
  var rec map[string]any
  json.NewDecoder(resp.Body).Decode(&rec)
  resp.Body.Close()
  path := fmt.Sprintf("%s/records/%v", srv.URL, rec["id"])
  get := intReq(t, "GET", path, "")
  if get.StatusCode != 200 {
    t.Fatalf("want 200 got %d", get.StatusCode)
  }
  get.Body.Close()
  del := intReq(t, "DELETE", path, "")
  if del.StatusCode != 200 {
    t.Fatalf("want 200 got %d", del.StatusCode)
  }
  del.Body.Close()
  gone := intReq(t, "GET", path, "")
  if gone.StatusCode != 404 {
    t.Fatalf("want 404 got %d", gone.StatusCode)
  }
  gone.Body.Close()
}
