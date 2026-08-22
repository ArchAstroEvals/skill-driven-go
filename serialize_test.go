package main

import ("testing"
  "reflect")

func TestSelectFields(t *testing.T) {
  rec := Record{"a": 1, "b": 2}
  if got := selectFields(rec, []string{"a"}); !reflect.DeepEqual(got, Record{"a": 1}) {
    t.Fatalf("want subset got %v", got)
  }
}
