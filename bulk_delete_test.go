package main

import "testing"

func TestBulkDelete(t *testing.T) {
  s := NewStore()
  s.Create(map[string]any{"name": "a"})
  s.Create(map[string]any{"name": "b"})
  s.Create(map[string]any{"name": "c"})
  if n := s.BulkDelete([]int{1, 2, 99}); n != 2 {
    t.Fatalf("want 2 deleted got %d", n)
  }
  if len(s.List()) != 1 {
    t.Fatal("want one left")
  }
}
