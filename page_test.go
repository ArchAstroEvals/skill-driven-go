package main

import ("testing"
  "reflect")

func TestPages(t *testing.T) {
  recs := []Record{{"id": 1}, {"id": 2}, {"id": 3}, {"id": 4}}
  if got := paginate(recs, 1, 3); len(got) != 3 {
    t.Fatal("want first page")
  }
  if got := paginate(recs, 2, 3); len(got) != 1 {
    t.Fatal("want second page")
  }
  _ = reflect.DeepEqual
}
