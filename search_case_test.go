package main

import "testing"

func TestSearchCase(t *testing.T) {
  recs := []Record{{"name": "Widget"}}
  if got := searchByName(recs, "widget"); len(got) != 1 {
    t.Fatal("want case-insensitive match")
  }
}
