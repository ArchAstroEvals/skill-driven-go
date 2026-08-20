package main

import "testing"

func TestFieldTypes(t *testing.T) {
  if fieldType("x") != "string" || fieldType(1.5) != "number" || fieldType(true) != "boolean" || fieldType(nil) != "unknown" {
    t.Fatal("bad type mapping")
  }
  if got := required(map[string]any{"name": ""}, []string{"name"}); len(got) != 1 {
    t.Fatal("want blank rejected")
  }
}
