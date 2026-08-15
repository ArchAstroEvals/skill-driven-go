package main

import "testing"

func TestStoreListDelete(t *testing.T) {
  s := NewStore()
  s.Create(map[string]any{"name": "a"})
  s.Create(map[string]any{"name": "b"})
  if len(s.List()) != 2 {
    t.Fatal("want two records")
  }
  if !s.Delete(1) {
    t.Fatal("want delete true")
  }
  if s.Delete(1) {
    t.Fatal("want second delete false")
  }
}
