package main

import ("encoding/csv"
  "strings"
  "testing")

func TestCSV(t *testing.T) {
  out := toCSV([]Record{{"a": "x", "b": "y"}}, []string{"a", "b"})
  rows, err := csv.NewReader(strings.NewReader(out)).ReadAll()
  if err != nil || len(rows) != 2 || rows[1][1] != "y" {
    t.Fatalf("bad csv %q", out)
  }
}
