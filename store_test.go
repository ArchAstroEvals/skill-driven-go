package main

import "testing"

func TestStoreCreateGet(t *testing.T) {
  s := NewStore()
  rec := s.Create(map[string]any{"name": "Widget"})
  if rec["id"] != 1 {
    t.Fatalf("want id 1 got %v", rec["id"])
  }
  got, ok := s.Get(1)
  if !ok || got["name"] != "Widget" {
    t.Fatal("want stored widget")
  }
  if _, ok := s.Get(99); ok {
    t.Fatal("want missing record")
  }
}
