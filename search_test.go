package main

import "testing"

func TestSearch(t *testing.T) {
  recs := []Record{{"name": "Widget"}, {"name": "Gadget"}}
  got := searchByName(recs, "Wid")
  if len(got) != 1 {
    t.Fatal("want one match")
  }
}
