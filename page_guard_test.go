package main

import "testing"

func TestPageZero(t *testing.T) {
  recs := []Record{{"id": 1}}
  if got := paginate(recs, 0, 3); len(got) != 0 {
    t.Fatal("want empty page zero")
  }
}
