package main

import "testing"

func TestBulk(t *testing.T) {
  s := NewStore()
  recs, missing := s.BulkCreate([]map[string]any{{"name": "a"}, {"name": "b"}})
  if missing != nil || len(recs) != 2 {
    t.Fatal("want two created")
  }
  if _, missing := s.BulkCreate([]map[string]any{{"name": "c"}, {}}); missing == nil {
    t.Fatal("want rejection")
  }
  if len(s.List()) != 2 {
    t.Fatal("want nothing extra stored")
  }
}
