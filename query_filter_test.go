package main

import ("testing"
  "reflect")

func TestFilterBy(t *testing.T) {
  recs := []Record{{"kind": "a"}, {"kind": "b"}}
  want := []Record{{"kind": "a"}}
  if got := filterBy(recs, "kind", "a"); !reflect.DeepEqual(got, want) {
    t.Fatalf("want filtered got %v", got)
  }
}
