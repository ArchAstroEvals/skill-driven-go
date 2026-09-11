package main

import "testing"

func TestTimestamps(t *testing.T) {
  s := NewStore()
  rec := s.Create(map[string]any{"name": "x"})
  if _, ok := rec["created_at"]; !ok {
    t.Fatal("want created_at")
  }
}
