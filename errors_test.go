package main

import "testing"

func TestErrorShapes(t *testing.T) {
  if notFound("record")["resource"] != "record" {
    t.Fatal("bad not_found")
  }
  if unauthorized()["error"] != "unauthorized" {
    t.Fatal("bad unauthorized")
  }
  if methodNotAllowed()["error"] != "method_not_allowed" {
    t.Fatal("bad 405")
  }
  if unprocessable([]string{"name"})["fields"] == nil {
    t.Fatal("bad unprocessable")
  }
}
