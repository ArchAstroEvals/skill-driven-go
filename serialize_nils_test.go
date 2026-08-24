package main

import ("testing"
  "reflect")

func TestDropNils(t *testing.T) {
  rec := Record{"a": nil, "b": 2}
  if got := selectFields(rec, []string{"a", "b"}); !reflect.DeepEqual(got, Record{"b": 2}) {
    t.Fatalf("want nils dropped got %v", got)
  }
}
