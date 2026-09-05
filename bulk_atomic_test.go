package main

import "testing"

func TestBulkAtomic(t *testing.T) {
  s := NewStore()
  if _, missing := s.BulkCreate([]map[string]any{{}, {"name": "b"}}); missing == nil {
    t.Fatal("want rejection")
  }
  if len(s.List()) != 0 {
    t.Fatal("want empty store")
  }
}
