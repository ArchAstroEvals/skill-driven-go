package main

import ("encoding/csv"
  "strings"
  "testing")

func TestCSVQuotes(t *testing.T) {
  out := toCSV([]Record{{"a": "x,y"}}, []string{"a"})
  rows, err := csv.NewReader(strings.NewReader(out)).ReadAll()
  if err != nil || rows[1][0] != "x,y" {
    t.Fatalf("bad round trip %q", out)
  }
}
