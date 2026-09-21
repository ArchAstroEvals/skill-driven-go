package main

import "testing"

func TestBulkEmpty(t *testing.T) {
  s := NewStore()
  recs, missing := s.BulkCreate([]map[string]any{})
  if missing != nil || len(recs) != 0 {
    t.Fatal("want empty success")
  }
}
