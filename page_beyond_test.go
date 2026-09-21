package main

import "testing"

func TestPageBeyond(t *testing.T) {
  recs := []Record{{"id": 1}, {"id": 2}}
  if got := paginate(recs, 99, 10); len(got) != 0 {
    t.Fatal("want empty far page")
  }
}
